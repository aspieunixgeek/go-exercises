package wtconv

import "fmt"

type Kg float64
type Lb float64

func (kg Kg) String() string {
	return fmt.Sprintf("%.2f Kg", kg)
}

func (lb Lb) String() string {
	return fmt.Sprintf("%.2f lb.", lb)
}
