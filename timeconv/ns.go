// Nanosecond converter

package timeconv

// nanosecond to second
func NsToSec(ns NanoSec) Sec {
	return Sec(ns / 1e+9)
}

// nanosecond to millisecond
func NsToMs(ns NanoSec) MilliSec {
	return MilliSec(ns / 1e+6)
}

// nanosecond to micro-second
func NsToToMicroSec(ns NanoSec) MicroSec {
	return MicroSec(ns / 1000)
}

func NsToPs(ns NanoSec) PicoSec {
	return PicoSec(ns * 1000)
}
