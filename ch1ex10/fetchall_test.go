package ch1ex10

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
	urls := []string{"https://youtube.com", "https://youtube.com"}

	time.Sleep(2000 * time.Millisecond)
	out, err := FetchAll(urls)
	if err != nil {
		t.Errorf("FetchAll: %v\n", err)
	}

	defer out.Close() // ignoring
	_, err = out.Read(b)
	t.Logf("%s\n", b)
}
