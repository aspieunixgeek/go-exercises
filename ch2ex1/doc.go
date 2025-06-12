// Exercise 2.1:
// Add types, constants, and functions to tempconv for processing temperatures in the Kelvin scale,
// where zero Kelvin is -213.15°C and a difference of 1K has the same magnitude as 1°C.

/*
# Init go mod

	$ go mod init ch2ex1

# Run unit tests

	$ go test -v
	=== RUN   Test_CToF
	    conv_test.go:10: 0.8312325142909469
	--- PASS: Test_CToF (0.00s)
	=== RUN   Test_CToK
	    conv_test.go:22: 0.520817797954284
	--- PASS: Test_CToK (0.00s)
	=== RUN   Test_KToC
	    conv_test.go:34: 0.5605857883336317
	--- PASS: Test_KToC (0.00s)
	=== RUN   Test_KToF
	    conv_test.go:46: 0.2798065191273187
	--- PASS: Test_KToF (0.00s)
	=== RUN   Test_FToC
	    conv_test.go:58: 0.2623660992285545
	--- PASS: Test_FToC (0.00s)
	=== RUN   Test_FToK
	    conv_test.go:70: 0.4738559653952945
	--- PASS: Test_FToK (0.00s)
	PASS
	ok      ch2ex1  0.002s
*/
package ch2ex1
