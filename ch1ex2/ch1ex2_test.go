package ch1ex2

import (
	"os"
	"testing"
)

func TestEcho1(t *testing.T) {
	os.Args = []string{"a", "b", "c"}
	out := Echo1()

	correct := "index: 0, value: a\nindex: 1, value: b\nindex: 2, value: c\n"
	if out != correct {
		t.Errorf("\"%s\" Not Equal \"%s\"\n", out, correct)
	}
}
func TestEcho2(t *testing.T) {
	os.Args = []string{"a", "b", "c"}
	out := Echo2()

	correct := "index: 0, value: a\nindex: 1, value: b\nindex: 2, value: c\n"
	if out != correct {
		t.Errorf("\"%s\" Not Equal \"%s\"\n", out, correct)
	}
}

func TestEcho3(t *testing.T) {
	os.Args = []string{"a", "b", "c"}
	out := Echo3()

	correct := "index: 0, value: a\nindex: 1, value: b\nindex: 2, value: c"
	if out != correct {
		t.Errorf("\"%s\" Not Equal \"%s\"\n", out, correct)
	}
}
