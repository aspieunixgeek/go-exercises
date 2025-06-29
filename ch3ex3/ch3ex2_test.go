package ch3ex3

import (
	"errors"
	"io/fs"
	"os"
	"testing"
)

func Test_MkDir(t *testing.T) {
	err := os.Mkdir("out", 0750)
	if err != nil && !errors.Is(err, fs.ErrExist) {
		t.Errorf("os.Mkdir: %v\n", err)
	}
}

func Test_ChDir(t *testing.T) {
	err := os.Chdir("out")
	if err != nil {
		t.Errorf("os.Chdir: %v\n", err)
	}
}

func Test_F(t *testing.T) {
	out := Gen("red", "blue")
	f, err := os.Create("out.svg")
	if err != nil {
		t.Errorf("os.Create: %v\n", err)
	}

	_, err = f.WriteString(out)
	if err != nil {
		t.Errorf("f.WriteString: %v\n", err)
	}
	f.Close() // ignoring
}
