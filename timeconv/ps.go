// Picosecond converter

package timeconv

// picosecond to nanosecond
func PsToNs(ps PicoSec) NanoSec {
	return NanoSec(ps / 1000)
}

// picosecond to microsecond
func PsToMicroSec(ps PicoSec) MicroSec {
	return MicroSec(ps / 1e+6) // ps / 1 000 000 (million)
}

// picosecond to millisecond
func PsToMs(ps PicoSec) MilliSec {
	return MilliSec(ps / 1e+9) // ps / 1 000 000 000 (billion)
}

// picosecond to second
func PsToSec(ps PicoSec) Sec {
	return Sec(ps / 1e+12) // ps / 1 000 000 000 000 (trillion)
}
