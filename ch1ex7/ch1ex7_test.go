package ch1ex7

import (
	"testing"
)

func TestFetch(t *testing.T) {
	urls := []string{"https://google.com", "https://docker.com"}

	for _, url := range urls {

		n, err := Fetch(url)
		if err != nil {
			t.Errorf("Fetch(%s): %v\n", url, err)
		}
		if n == 0 {
			t.Errorf("empty response\n")
		}
	}
}
