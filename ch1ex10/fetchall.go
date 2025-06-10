package ch1ex10

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func FetchAll(urls []string) (*os.File, error) {
	var out string

	start := time.Now()
	ch := make(chan string)
	for i, url := range urls {
		go fetch(url, ch, i) // start a goroutine
	}
	for range urls {
		out += <-ch // receive from channel ch
	}

	out += fmt.Sprintf("%.3fs elapsed\n", time.Since(start).Seconds())

	now := time.Now().Format("02_01_2006__15_04_05")
	filename := fmt.Sprintf("out_%s.txt", now)
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close() // ignoring

	_, err = f.WriteString(out)

	if err != nil {
		return nil, err
	}
	return f, nil
}

func fetch(url string, ch chan<- string, index int) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.3fs %7d %s\n", secs, nbytes, url)
}
