package timeconv

import (
	"math/rand"
	"testing"
)

// second to nanosecond
func Test_SecToNs(t *testing.T) {
	r := rand.Float64()

	correct := NanoSec(r * 1e+9)
	needCheck := SecToNs(Sec(r))

	if needCheck != correct {
		t.Errorf("needCheck(%v) Not Equal correct(%v)", needCheck, correct)
	}
}

// second to millisecond
func Test_SecToMilliSec(t *testing.T) {
	r := rand.Float64()

	correct := MilliSec(r * 1000)
	needCheck := SecToMilliSec(Sec(r))

	if needCheck != correct {
		t.Errorf("needCheck(%v) Not Equal correct(%v)", needCheck, correct)
	}
}

// second to micro-second
func Test_SecToMicroSec(t *testing.T) {
	r := rand.Float64()

	correct := MicroSec(r * 1e+6)
	needCheck := SecToMicroSec(Sec(r))

	if needCheck != correct {
		t.Errorf("needCheck(%v) Not Equal correct(%v)", needCheck, correct)
	}
}
