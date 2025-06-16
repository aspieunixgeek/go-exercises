// Microsecond converter

package timeconv

// micro-second to nanosecond
func MicroSecToNanoSec(mS MicroSec) NanoSec {
	return NanoSec(mS * 1000)
}

// micro-second to millisecond
func MicroSecToMs(mS MicroSec) MilliSec {
	return MilliSec(mS / 1000)
}

// micro-second to second
func MicroSecToSec(mS MicroSec) Sec {
	return Sec(mS / 1e+6)
}
