package ch2ex1

import (
	"math/rand"
	"testing"
)

func Test_CToF(t *testing.T) {
	r := rand.Float64()
	t.Logf("%v", r)

	correct := Fahrenheit((r * 9 / 5) + 32)
	needCheck := CToF(Celsius(r))

	if correct != needCheck {
		t.Errorf("\"%v\" Not Equal \"%v\"", needCheck, correct)
	}
}

func Test_CToK(t *testing.T) {
	r := rand.Float64()
	t.Logf("%v", r)

	correct := Kelvin(r + 273.15)
	needCheck := CToK(Celsius(r))

	if correct != needCheck {
		t.Errorf("\"%v\" Not Equal \"%v\"", needCheck, correct)
	}
}

func Test_KToC(t *testing.T) {
	r := rand.Float64()
	t.Logf("%v", r)

	correct := Celsius(r - 273.15)
	needCheck := KToC(Kelvin(r))

	if correct != needCheck {
		t.Errorf("\"%v\" Not Equal \"%v\"", needCheck, correct)
	}
}

func Test_KToF(t *testing.T) {
	r := rand.Float64()
	t.Logf("%v", r)

	correct := Fahrenheit((r-273.15)*9/5 + 32)
	needCheck := KToF(Kelvin(r))

	if correct != needCheck {
		t.Errorf("\"%v\" Not Equal \"%v\"", needCheck, correct)
	}
}

func Test_FToC(t *testing.T) {
	r := rand.Float64()
	t.Logf("%v", r)

	correct := Celsius((r - 32) * 5 / 9)
	needCheck := FToC(Fahrenheit(r))

	if correct != needCheck {
		t.Errorf("\"%v\" Not Equal \"%v\"", needCheck, correct)
	}
}

func Test_FToK(t *testing.T) {
	r := rand.Float64()
	t.Logf("%v", r)

	correct := Kelvin((r-32)*5/9 + 273.15)
	needCheck := FToK(Fahrenheit(r))

	if correct != needCheck {
		t.Errorf("\"%v\" Not Equal \"%v\"", needCheck, correct)
	}
}
