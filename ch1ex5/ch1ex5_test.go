package ch1ex5

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestLissajous(t *testing.T) {
	err := os.Mkdir("out", 0750)
	if err != nil && !errors.Is(err, os.ErrExist) {
		t.Errorf("os.Mkdir: %v\n", err)
	}

	err = os.Chdir("out")
	if err != nil {
		t.Errorf("os.Chdir: %v\n", err)
	}

	filename := fmt.Sprintf("file_%d.gif", uint16(rand.Int()))

	f, err := os.Create(filename)
	if err != nil {
		t.Errorf("os.Create: %v\n", err)
	}
	Lissajous(f)
}
