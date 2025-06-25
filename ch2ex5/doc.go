// Exercise 2.5:
// The expression x&(x-1) clears the rightmost non-zero bit of x.
// Write a version of PopCount that counts bits by using this fact,
// and assess its performance.
/*

# Unit tests

	$ go test -v
	=== RUN   Test_CompareBitsCount
	    ch2ex5_test.go:18: PopCountNew: 39
	    ch2ex5_test.go:19: PopCountOld: 39
	--- PASS: Test_CompareBitsCount (0.00s)
	PASS
	ok      github.com/aspieunixgeek/go-exercises/ch2ex5    0.002s

*/
package ch2ex5
