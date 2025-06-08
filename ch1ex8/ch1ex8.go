package ch1ex8

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
		return nil, fmt.Errorf("fetch: get %v\n", err)
	}
	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("fetch: read %s %v\n", url, err)
	}

	return b, nil
}
