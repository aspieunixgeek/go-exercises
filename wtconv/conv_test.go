package wtconv

import (
	"math/rand"
	"testing"
)

func Test_KgToLb(t *testing.T) {
	r := rand.Float64()

	correct := Lb(r * 2.205)
	needCheck := KgToLb(Kg(r))

	if correct != needCheck {
		t.Errorf("\"%v\" Not Equal \"%v\"", needCheck, correct)
	}
}

func Test_LbToKg(t *testing.T) {
	r := rand.Float64()

	correct := Kg(r / 2.205)
	needCheck := LbToKg(Lb(r))

	if correct != needCheck {
		t.Errorf("\"%v\" Not Equal \"%v\"", needCheck, correct)
	}
}
