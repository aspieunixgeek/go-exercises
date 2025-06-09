package ch1ex9

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Fetch(url string) ([]byte, error) {

	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetch: %v\n", err)
	}
	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("fetch: reading %s: %v\n", url, err)
	}

	status := fmt.Sprintf("<!-- %s -->", resp.Status)
	bb := []byte(status)
	b = append(b, bb...)
	return b, nil
}
