package ch1ex11

import (
	"errors"
	"os"
	"testing"
	"time"
)

func TestCreateOutDir(t *testing.T) {
	err := os.Mkdir("out", 0750)
	if err != nil && !errors.Is(err, os.ErrExist) {
		t.Errorf("os.Mkdir: %v\n", err)
	}

	err = os.Chdir("out")
	if err != nil {
		t.Errorf("os.Chdir: %v\n", err)
	}
}

func TestFetchAll(t *testing.T) {
	var b = make([]byte, 1024)
	urls := []string{
		"https://youtube.com", "https://www.dw.com", "https://www.bbc.com", "https://docs.docker.com", "https://whatsapp.com",
		"https://stackoverflow.com", "https://docs.docker.com", "https://github.com", "https://www.reddit.com", "https://ubuntu.com", "https://amazon.com",
		"https://wikipedia.org", "https://yahoo.com", "https://live.com", "https://office.com", "https://samsung.com", "https://duckduckgo.com", "https://aliexpress.com",
		"https://roblox.com", "https://zoom.us", "https://notwork.com",
	}

	out, err := FetchAll(urls, 8*time.Second) // request-timeout = 8 sec.
	if err != nil {
		t.Errorf("FetchAll: %v\n", err)
	}

	defer out.Close() // ignoring
	_, err = out.Read(b)
	t.Logf("%s\n", b)
}
