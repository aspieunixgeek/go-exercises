package ch3ex1

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"
)

// create a folder and enter it
func Test_ChDir(t *testing.T) {
	err := os.Mkdir("out", 0750)
	if err != nil && !errors.Is(err, os.ErrExist) { // ignoring exist dir
		t.Errorf("os.Mkdir: %v\n", err)
	}

	err = os.Chdir("out")
	if err != nil {
		t.Errorf("os.Chdir: %v\n", err)
	}
}

// generate SVG file
func Test_Gen(t *testing.T) {
	data := Gen()
	time.Sleep(1000 * time.Millisecond)

	n := time.Now().Format("02_01_2006__15_04_05")

	fl, err := os.Create(fmt.Sprintf("out_%s.svg", n))
	if err != nil {
		t.Errorf("os.Create: %v\n", err)
	}

	if len(data) == 0 {
		t.Errorf("len(data) == 0\n")
	}

	_, err = fl.WriteString(data)
	if err != nil {
		t.Errorf("fe.WriteString: %v\n", err)
	}
	fl.Close() // ignoring
}
