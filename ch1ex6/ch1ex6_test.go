package ch1ex6

import (
	"errors"
	"os"
	"testing"
)

func TestGen(t *testing.T) {

	wd, err := os.Getwd()
	if err != nil {
		t.Errorf("os.Getwd: %v", err)
	}

	if wd != "out" {

		err := os.Mkdir("out", 0750)
		if err != nil && !errors.Is(err, os.ErrExist) {
			t.Errorf("os.Mkdir: %v\n", err)
		}
		err = os.Chdir("out")
		if err != nil {
			t.Errorf("os.Chdir: %v\n", err)
		}
	}

	err = Gen()
	if err != nil {
		t.Errorf("Gen: %v\n", err)
	}
}
