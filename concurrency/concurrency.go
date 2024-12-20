// Package concurrency provides functions for checking websites concurrently.
package concurrency

// WebsiteChecker is a function type that takes a URL as input and returns a boolean
// indicating whether the website is valid or not.
type WebsiteChecker func(string) bool

// result is a struct used to hold the URL and the result of the WebsiteChecker for that URL.
type result struct {
	string
	bool
}

// CheckWebsites takes a WebsiteChecker and a slice of URLs, and returns a map of URLs to
// boolean values indicating whether each website is valid or not. It uses concurrency to
// check the websites in parallel.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(url string) {
			resultChannel <- result{url, wc(url)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
