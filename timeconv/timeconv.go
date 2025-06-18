package timeconv

import (
	"fmt"
)

type PicoSec float64
type NanoSec float64
type MicroSec float64
type MilliSec float64
type Sec float64

func (ps PicoSec) String() string {
	return fmt.Sprintf("%.12f ps", ps)
}

func (ns NanoSec) String() string {
	return fmt.Sprintf("%.12f ns", ns)
}

func (us MicroSec) String() string {
	return fmt.Sprintf("%.12f μs", us)
}

func (ms MilliSec) String() string {
	return fmt.Sprintf("%.12f ms", ms)
}

func (s Sec) String() string {
	return fmt.Sprintf("%.12f s", s)
}
