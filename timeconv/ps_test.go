package timeconv

import (
	"math/rand"
	"testing"
)

// picosecond to nanosecond
func Test_PsToNs(t *testing.T) {
	r := rand.Float64()

	correct := NanoSec(r / 1000)
	needCheck := PsToNs(PicoSec(r))

	if needCheck != correct {
		t.Errorf("needCheck(%v) Not Equal correct(%v)", needCheck, correct)
	}
	t.Logf("needCheck(%v), correct(%v)", needCheck, correct)
}

// picosecond tp millisecond
func Test_PsToMs(t *testing.T) {
	r := rand.Float64()

	correct := MilliSec(r / 1e+9) // r / 1 000 000 000
	needCheck := PsToMs(PicoSec(r))

	if needCheck != correct {
		t.Errorf("needCheck(%v) Not Equal correct(%v)", needCheck, correct)
	}
	t.Logf("needCheck(%v), correct(%v)", needCheck, correct)
}

// picosecond tp microsecond
func Test_PsToMicroSec(t *testing.T) {
	r := rand.Float64()

	correct := MicroSec(r / 1e+6) // r / 1 000 000
	needCheck := PsToMicroSec(PicoSec(r))

	if needCheck != correct {
		t.Errorf("needCheck(%v) Not Equal correct(%v)", needCheck, correct)
	}
	t.Logf("needCheck(%v), correct(%v)", needCheck, correct)
}

// picosecond tp second
func Test_PsToSec(t *testing.T) {
	r := rand.Float64()

	correct := Sec(r / 1e+12) // r / 1 000 000 000 000
	needCheck := PsToSec(PicoSec(r))

	if needCheck != correct {
		t.Errorf("needCheck(%v) Not Equal correct(%v)", needCheck, correct)
	}
	t.Logf("needCheck(%v), correct(%v)", needCheck, correct)
}
