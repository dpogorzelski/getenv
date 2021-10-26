package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	prefix := flag.String("prefix", "", "a string")
	flag.Parse()

	var c = &http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest("GET", "http://metadata/computeMetadata/v1/instance/attributes?recursive=true", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Metadata-Flavor", "Google")
	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatal("couldn't fetch metadata attributes")
	}

	items := make(map[string]string)
	vars := make(map[string]string)

	err = json.NewDecoder(res.Body).Decode(&items)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range items {
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
