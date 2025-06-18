package timeconv

import (
	"math/rand"
	"testing"
)

// nanosecond to second
func Test_NsToSec(t *testing.T) {
	r := rand.Float64()

	correct := Sec(r / 1e+9) // r / 1 000 000 000
	needCheck := NsToSec(NanoSec(r))

	if needCheck != correct {
		t.Errorf("needCheck(%v) Not Equal correct(%v)", needCheck, correct)
	}
	t.Logf("needCheck(%v), correct(%v)", needCheck, correct)
}

// nanosecond to milliSecond
func Test_NsToMs(t *testing.T) {
	r := rand.Float64()

	correct := MilliSec(r / 1e+6) // r / 1 000 000
	needCheck := NsToMs(NanoSec(r))

	if needCheck != correct {
		t.Errorf("needCheck(%v) Not Equal correct(%v)", needCheck, correct)
	}
	t.Logf("needCheck(%v), correct(%v)", needCheck, correct)
}

// nanosecond to micro-second
func Test_NsToMicroSec(t *testing.T) {
	r := rand.Float64()

	correct := MicroSec(r / 1000)
	needCheck := NsToToMicroSec(NanoSec(r))

	if needCheck != correct {
		t.Errorf("needCheck(%v) Not Equal correct(%v)", needCheck, correct)
	}
	t.Logf("needCheck(%v), correct(%v)", needCheck, correct)
}

// nanosecond to picosecond
func Test_NsToPs(t *testing.T) {
	r := rand.Float64()

	correct := PicoSec(r * 1000)
	needCheck := NsToPs(NanoSec(r))

	if needCheck != correct {
		t.Errorf("needCheck(%v) Not Equal correct(%v)", needCheck, correct)
	}
	t.Logf("needCheck(%v), correct(%v)", needCheck, correct)
}
