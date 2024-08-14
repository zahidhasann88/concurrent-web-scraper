package scraper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// FetchURL fetches data from a URL and sends the result to a channel
func FetchURL(url string, ch chan<- Result) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- Result{URL: url, Error: fmt.Errorf("error fetching %s: %v", url, err)}
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- Result{URL: url, Error: fmt.Errorf("error reading response from %s: %v", url, err)}
		return
	}

	duration := time.Since(start).Seconds()
	ch <- Result{URL: url, Duration: duration, Size: len(body)}
}

// FetchURLsConcurrently fetches multiple URLs concurrently and returns the results
func FetchURLsConcurrently(urls []string) []Result {
	ch := make(chan Result, len(urls))
	results := make([]Result, len(urls))

	for _, url := range urls {
		go FetchURL(url, ch)
	}

	for i := range urls {
		results[i] = <-ch
	}

	return results
}
