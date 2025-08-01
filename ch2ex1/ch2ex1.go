package ch2ex1

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius    = -273.15
	AbsoluteZeroK Kelvin     = 0
	AbsoluteZeroF Fahrenheit = -459.67

	FreezingC Celsius    = 0
	FreezingK Kelvin     = 273.15
	FreezingF Fahrenheit = 32

	BoilingC Celsius    = 100
	BoilingK Kelvin     = 373.15
	BoilingF Fahrenheit = 212
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}
