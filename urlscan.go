package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type Result struct {
	Task struct {
		URL string `json:"url"`
	} `json:"page"`
}

func main() {
	// define command line flag to specify the domain name
	domainPtr := flag.String("url", "docs.google.com", "the domain name to search for")
	flag.Parse()

	// construct the URL with the specified domain name
	url := fmt.Sprintf("https://urlscan.io/api/v1/search/?q=domain:%s&size=1000", *domainPtr)

	// perform the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// parse the response body as JSON
	var data struct {
		Results []Result `json:"results"`
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	// store unique URLs in a map
	urls := make(map[string]bool)
	for _, result := range data.Results {
		urls[result.Task.URL] = true
	}

	// print the unique URLs to the command line
	for url := range urls {
		fmt.Println(url)
	}
}
