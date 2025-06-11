package ch1ex11

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func FetchAll(urls []string, timeout time.Duration) (*os.File, error) {
	var out string

	start := time.Now()
	ch := make(chan string)

	for _, url := range urls {
		go fetch(url, ch, timeout) // start a goroutine
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

func fetch(url string, ch chan<- string, reqTimeout time.Duration) {
	start := time.Now()

	client := &http.Client{Timeout: reqTimeout}
	resp, err := client.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%.3fs %v\n", time.Since(start).Seconds(), err) // send to channel ch
		return                                                            //  quit signal for current go-routine
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return //  quit signal for current go-routine
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.3fs %7d %s\n", secs, nbytes, url)
}
