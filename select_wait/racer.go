// Package select_wait implements a function that returns the URL of the
// server that responds first.
package select_wait

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

// Racer takes two urls and returns the one that responds first, or an error
// if the timeout is reached before either responds.
func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// ConfigurableRacer takes two URLs and a timeout duration, returning the URL
// of the server that responds first. If neither server responds within the
// specified timeout, an error is returned indicating a timeout occurred.

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// ping starts a goroutine that sends a GET request to the specified URL
// and then closes the returned channel. This allows the caller to wait
// for the request to complete using a select statement.
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
