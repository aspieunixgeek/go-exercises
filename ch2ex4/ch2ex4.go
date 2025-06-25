package ch2ex4

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountOld returns the populations count (number of set bits) of x.
func PopCountOld(x uint64) int {
	y := int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))] +
		pc[byte(x>>(8*8-1))])

	return y
}

func PopCountNew(x uint64) int {
	var y byte

	for i := 0; i <= 8; i++ {
		if i == 8 {
			y += pc[byte(x>>(i*8-1))]
		} else {
			y += pc[byte(x>>(i*8))]
		}
	}

	return int(y)
}
