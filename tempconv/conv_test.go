package tempconv

import (
	"math/rand"
	"testing"
)

// celsius to fahrenheit
func Test_CToF(t *testing.T) {
	r := rand.Float64()

	correct := Fahrenheit((r * 9 / 5) + 32)
	needCheck := CToF(Celsius(r))

	if correct != needCheck {
		t.Errorf("\"%v\" Not Equal \"%v\"", needCheck, correct)
	}
}

// fahrenheit to celsius.
func Test_FToC(t *testing.T) {
	r := rand.Float64()

	correct := Celsius((r - 32) * 5 / 9)
	needCheck := FToC(Fahrenheit(r))

	if correct != needCheck {
		t.Errorf("\"%v\" Not Equal \"%v\"", needCheck, correct)
	}
}
