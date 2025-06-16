// Millisecond converter

package timeconv

// millisecond to second
func MsToSec(ms MilliSec) Sec {
	return Sec(ms / 1000)
}

// millisecond to nanosecond
func MsToNs(ms MilliSec) NanoSec {
	return NanoSec(ms * 1e+6)
}

// millisecond to micro-second
func MsToMicroSec(ms MilliSec) MicroSec {
	return MicroSec(ms * 1000)
}
