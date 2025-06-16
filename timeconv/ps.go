// Picosecond converter

package timeconv

// Picosecond to Nanosecond
func PsToNs(ps PicoSec) NanoSec {
	return NanoSec(ps / 1000)
}

// Picosecond to Microsecond
func PsToMicSec(ps PicoSec) MicroSec {
	return MicroSec(ps / 1e+6) // ps / 1 000 000 (million)
}

// Picosecond to Millisecond
func PsToMs(ps PicoSec) MilliSec {
	return MilliSec(ps / 1e+9) // ps / 1 000 000 000 (billion)
}

// Picosecond to Second
func PsToSec(ps PicoSec) Sec {
	return Sec(ps / 1e+12) // ps / 1 000 000 000 000 (trillion)
}
