package ch2ex5

import (
	"math/rand"
	"testing"
)

func Test_CompareBitsCount(t *testing.T) {
	var x = rand.Uint64()

	cntNew := PopCountNew(x)
	cntOld := PopCountOld(x)

	if cntNew != cntOld {
		t.Errorf("PopCountNew(%v) Not Equal PopCountOld(%v)\n", cntNew, cntOld)
	}

	t.Logf("PopCountNew: %v\n", cntNew)
	t.Logf("PopCountOld: %v\n", cntOld)
}
