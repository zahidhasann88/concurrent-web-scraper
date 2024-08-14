package main

import (
	"fmt"
	"log"

	"concurrent-web-scraper/scraper"
)

func main() {
	urls := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://gophercises.com",
	}

	results := scraper.FetchURLsConcurrently(urls)

	for _, result := range results {
		if result.Error != nil {
			log.Printf("Failed to fetch %s: %v\n", result.URL, result.Error)
			continue
		}
		fmt.Printf("Fetched %s in %.2f seconds (%d bytes)\n", result.URL, result.Duration, result.Size)
	}
}
