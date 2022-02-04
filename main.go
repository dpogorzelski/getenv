package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	prefix := flag.String("prefix", "", "Only fetch metadata items that match this prefix")
	prepend := flag.String("prepend", "", "Prepend this string to the ENV var names in the output")
	flag.Parse()

	var client = &http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest("GET", "http://metadata/computeMetadata/v1/instance/attributes?recursive=true", nil)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	req.Header.Add("Metadata-Flavor", "Google")
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Fprint(os.Stderr, "couldn't fetch metadata attributes")
		os.Exit(1)
	}

	items := make(map[string]string)
	vars := make(map[string]string)

	err = json.NewDecoder(res.Body).Decode(&items)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	for k, v := range items {
		if strings.HasPrefix(k, *prefix) {
			vars[strings.ToUpper(k)] = v
		}
	}

	for k, v := range vars {
		if *prepend != "" {
			fmt.Printf("%s_%s=%s\n", strings.ToUpper(*prepend), k, v)
		} else {
			fmt.Printf("%s=%s\n", k, v)
		}
	}
}
