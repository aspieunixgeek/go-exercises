package ch1ex8

import (
	"testing"
)

// TestFetch1
// testing without http protocol
func TestFetch1(t *testing.T) {
	url := "docker.com"

	b, err := Fetch(url)
	if err != nil {
		t.Errorf("Fetch: %v\n", err)
	}

	if len(b) == 0 {
		t.Error("len(b) == 0\n")
	}
}

// TestFetch2
// testing with http protocol
func TestFetch2(t *testing.T) {
	url := "http://docker.com"

	b, err := Fetch(url)
	if err != nil {
		t.Errorf("Fetch: %v\n", err)
	}

	if len(b) == 0 {
		t.Error("len(b) == 0\n")
	}
}

// TestFetch3
// testing with https protocol
func TestFetch3(t *testing.T) {
	url := "https://docker.com"

	b, err := Fetch(url)
	if err != nil {
		t.Errorf("Fetch: %v\n", err)
	}

	if len(b) == 0 {
		t.Error("len(b) == 0\n")
	}
}
