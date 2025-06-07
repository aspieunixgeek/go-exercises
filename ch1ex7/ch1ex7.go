package ch1ex7

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Fetch(url string) (int64, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("http.Get: %v\n", err)
	}

	n, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close() // ignoring
	if err != nil {
		return 0, fmt.Errorf("io.Copy:  %s: %v\n", url, err)
	}

	return n, nil
}
