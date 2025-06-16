// Seconds converter

package timeconv

// second to nanosecond
func SecToNs(s Sec) NanoSec {
	return NanoSec(s * 1e+9)
}

// second to micro-second
func SecToMicroSec(s Sec) MicroSec {
	return MicroSec(s * 1e+6)
}

// second to millisecond
func SecToMilliSec(s Sec) MilliSec {
	return MilliSec(s * 1000)
}
