package timeconv

import (
	"math/rand"
	"testing"
)

// millisecond to second
func Test_MsToSec(t *testing.T) {
	r := rand.Float64()

	correct := Sec(r / 1000)
	needCheck := MsToSec(MilliSec(r))

	if needCheck != correct {
		t.Errorf("needCheck(%v) Not Equal correct(%v)", needCheck, correct)
	}
}

// millisecond to nanosecond
func Test_MsToNs(t *testing.T) {
	r := rand.Float64()

	correct := NanoSec(r * 1e+6)
	needCheck := MsToNs(MilliSec(r))

	if needCheck != correct {
		t.Errorf("needCheck(%v) Not Equal correct(%v)", needCheck, correct)
	}
}

// millisecond to micro-second
func Test_MsToMicroSec(t *testing.T) {
	r := rand.Float64()

	correct := MicroSec(r * 1000)
	needCheck := MsToMicroSec(MilliSec(r))

	if needCheck != correct {
		t.Errorf("needCheck(%v) Not Equal correct(%v)", needCheck, correct)
	}
}
