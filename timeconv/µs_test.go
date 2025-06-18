package timeconv

import (
	"math/rand"
	"testing"
)

func Test_MicroSecToNanoSec(t *testing.T) {
	r := rand.Float64()

	correct := NanoSec(r * 1000)
	needCheck := MicroSecToNanoSec(MicroSec(r))

	if needCheck != correct {
		t.Errorf("needCheck(%v) Not Equal correct(%v)", needCheck, correct)
	}
	t.Logf("needCheck(%v), correct(%v)", needCheck, correct)
}

func Test_MicroSecToMs(t *testing.T) {
	r := rand.Float64()

	correct := MilliSec(r / 1000)
	needCheck := MicroSecToMs(MicroSec(r))

	if needCheck != correct {
		t.Errorf("needCheck(%v) Not Equal correct(%v)", needCheck, correct)
	}
	t.Logf("needCheck(%v), correct(%v)", needCheck, correct)
}

func Test_MicroSecToSec(t *testing.T) {
	r := rand.Float64()

	correct := Sec(r / 1e+6)
	needCheck := MicroSecToSec(MicroSec(r))

	if needCheck != correct {
		t.Errorf("needCheck(%v) Not Equal correct(%v)", needCheck, correct)
	}
	t.Logf("needCheck(%v), correct(%v)", needCheck, correct)
}
