package lenconv

import "fmt"

type Meter float64
type Foot float64

func (m Meter) String() string {
	return fmt.Sprintf("%0.2f M", m)
}

func (f Foot) String() string {
	return fmt.Sprintf("%0.2f ft.", f)
}
