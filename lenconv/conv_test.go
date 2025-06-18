package lenconv

import (
	"math/rand"
	"testing"
)

func Test_FtToMet(t *testing.T) {
	r := rand.Float64()

	correct := Meter(r / 3.281)
	needCheck := FtToMet(Foot(r))

	if needCheck != correct {
		t.Errorf("correct: \"%v\" Not Equal needCheck: \"%v\"", needCheck, correct)
	}
}

func Test_MetToFt(t *testing.T) {
	r := rand.Float64()

	correct := Foot(r * 3.281)
	needCheck := MetToFt(Meter(r))

	if needCheck != correct {
		t.Errorf("correct: \"%v\" Not Equal needCheck: \"%v\"", correct, needCheck)
	}
}
