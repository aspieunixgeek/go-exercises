package ch2ex3

import (
	"math/rand"
	"testing"
)

func TestPopCountOld(t *testing.T) {
	const dec = 20 // bin: 10100

	cnt := PopCountOld(dec)

	if cnt != 2 {
		t.Errorf("cnt != 2")
	}
}

func TestPopCountNew(t *testing.T) {
	const dec = 20 // bin: 10100

	cnt := PopCountNew(dec)

	if cnt != 2 {
		t.Errorf("cnt != 4")
	}
}

func TestCompare(t *testing.T) {
	oldVersion := PopCountOld(20)
	newVersion := PopCountNew(20)

	if newVersion != oldVersion {
		t.Errorf("newVersion: \"%v\" Not Equal oldVersion: \"%v\"", newVersion, oldVersion)
	}
}

func BenchmarkPopCountOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountOld(rand.Uint64())
	}
}

func BenchmarkPopCountNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountNew(rand.Uint64())
	}
}
