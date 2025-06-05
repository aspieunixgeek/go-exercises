package ch1ex4

import (
	"os"
	"testing"
)

func TestDup2(t *testing.T) {
	err := os.Chdir("testdata")
	if err != nil {
		t.Errorf("os.Chdir: %v\n", err)
	}

	appName := os.Args[0]
	os.Args = []string{appName, "names1", "names2", "names3"}
	f := Dup2()
	for fileName, cnt := range f {
		if fileName == "names1/Alan" && cnt != 2 {
			t.Errorf("file: \"%s\", cnt Not Equal 3", fileName)
		}
		if fileName == "names2/Aidan" && cnt != 3 {
			t.Errorf("file: \"%s\", cnt Not Equal: %d\n", fileName, cnt)
		}
		if fileName == "names3/Dias" && cnt != 4 {
			t.Errorf("file: \"%s\", cnt Not Equal: %d\n", fileName, cnt)
		}
	}
}
