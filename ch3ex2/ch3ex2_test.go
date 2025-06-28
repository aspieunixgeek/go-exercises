package ch3ex2

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

func Test_F1(t *testing.T) {
	out := Gen(f1)
	f, err := os.Create("out_f1.svg")
	if err != nil {
		t.Errorf("os.Create: %v\n", err)
	}

	_, err = f.WriteString(out)
	if err != nil {
		t.Errorf("f.WriteString: %v\n", err)
	}
	f.Close() // ignoring
}

func Test_F2(t *testing.T) {
	out := Gen(f2)
	f, err := os.Create("out_f2.svg")
	if err != nil {
		t.Errorf("os.Create: %v\n", err)
	}

	_, err = f.WriteString(out)
	if err != nil {
		t.Errorf("f.WriteString: %v\n", err)
	}
	f.Close() // ignoring
}

func Test_F3(t *testing.T) {
	out := Gen(f3)
	f, err := os.Create("out_f3.svg")
	if err != nil {
		t.Errorf("os.Create: %v\n", err)
	}

	_, err = f.WriteString(out)
	if err != nil {
		t.Errorf("f.WriteString: %v\n", err)
	}
	f.Close() // ignoring
}

func Test_F4(t *testing.T) {
	out := Gen(f4)
	f, err := os.Create("out_f4.svg")
	if err != nil {
		t.Errorf("os.Create: %v\n", err)
	}

	_, err = f.WriteString(out)
	if err != nil {
		t.Errorf("f.WriteString: %v\n", err)
	}
	f.Close() // ignoring
}

func Test_F5(t *testing.T) {
	out := Gen(f5)
	f, err := os.Create("out_f5.svg")
	if err != nil {
		t.Errorf("os.Create: %v\n", err)
	}

	_, err = f.WriteString(out)
	if err != nil {
		t.Errorf("f.WriteString: %v\n", err)
	}
	f.Close() // ignoring
}
