package ch1ex9

import (
	"strings"
	"testing"
)

func TestFetch(t *testing.T) {
	url := "docker.com"

	out, err := Fetch(url)
	if err != nil {
		t.Errorf("Fetch: %s %v\n", url, err)
	}

	s := strings.TrimSpace(string(out[len(out)-16:])) // check tail of response
	t.Logf("%s\n", s)
	if s != "<!-- 200 OK -->" {
		t.Errorf("Error \"<!-- 200 OK -->\" Not Found\n")
	}
}
