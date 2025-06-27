package ch2ex5

import (
	"math/rand"
	"testing"
)

var x = rand.Uint64()

func Test_CompareBitsCount(t *testing.T) {
	cntNew := PopCountNew(x)
	cntOld := PopCountOld(x)

	if cntNew != cntOld {
		t.Errorf("PopCountNew(%d) Not Equal PopCountOld(%d)\n", cntNew, cntOld)
	}

	t.Logf("PopCountNew: bin(%[1]b), dec(%[1]d)\n", cntNew)
	t.Logf("PopCountOld: bin(%[1]b), dec(%[1]d)\n", cntOld)
}

func BenchmarkOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountOld(x)
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountNew(x)
	}
}
