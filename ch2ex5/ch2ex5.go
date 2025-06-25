package ch2ex5

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountNew(x uint64) int {
	var cnt int

	for {

		if x > 0 {
			cnt++
		} else {
			break
		}

		x = x & (x - 1)
	}

	return cnt
}

// PopCountOld returns the populations count (number of set bits) of x.
func PopCountOld(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])

}
