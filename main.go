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
	prefix := flag.String("prefix", "", "a string")
	flag.Parse()
	// keys := flag.Args()

	var c = &http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest("GET", "http://metadata/computeMetadata/v1/instance/attributes?recursive=true", nil)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	req.Header.Add("Metadata-Flavor", "Google")
	res, err := c.Do(req)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	defer res.Body.Close()
	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Fprint(os.Stderr, err)
	// 	os.Exit(1)
	// }

	if res.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "couldn't fetch metadata attributes")
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
		// matched, err := regexp.MatchString(`a.b`, "aaxbb")
		// if err != nil {
		// 	fmt.Fprint(os.Stderr, err)
		// 	os.Exit(1)
		// }
		vars[strings.ToUpper(k)] = v
	}

	for k, v := range vars {
		if *prefix != "" {
			fmt.Printf("%s_%s=%s\n", *prefix, k, v)
		} else {
			fmt.Printf("%s=%s\n", k, v)
		}
	}
}
