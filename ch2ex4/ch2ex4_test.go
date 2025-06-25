package ch2ex4

import (
	"math/rand"
	"testing"
)

var x = rand.Uint64()

// PopCount returns the populations count (number of set bits) of x.
func Test_PopCount(t *testing.T) {

	t.Logf("x: %v\n", x)
	if PopCountNew(x) != PopCountOld(x) {
		t.Errorf("\"PopCountNew(%v)\" Not Equal \"PopCountOld(%v)\"\n", PopCountNew(x), PopCountOld(x))
	}
}

func Benchmark_PopCountOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountOld(x)
	}
}

func Benchmark_PopCountNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountNew(x)
	}
}
