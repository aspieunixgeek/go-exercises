// Exercise 1.1
// Modify the echo program to also print os.Args[0],
// the name of the command that invoked it.

/*

# Run tests without flags

	$ go test
	/tmp/go-build3630873470/b001/ch1ex1.test -test.paniconexit0 -test.timeout=10m0s
	/tmp/go-build3630873470/b001/ch1ex1.test -test.paniconexit0 -test.timeout=10m0s
	/tmp/go-build3630873470/b001/ch1ex1.test -test.paniconexit0 -test.timeout=10m0s
	PASS
	ok      github.com/aspieunixgeek/go-exercises/ch1ex1    0.002s


# Run tests with v flag

	$ go test  -v
	=== RUN   TestEcho1
	/tmp/go-build1718110061/b001/ch1ex1.test -test.paniconexit0 -test.timeout=10m0s -test.v=true
	--- PASS: TestEcho1 (0.00s)
	=== RUN   TestEcho2
	/tmp/go-build1718110061/b001/ch1ex1.test -test.paniconexit0 -test.timeout=10m0s -test.v=true
	--- PASS: TestEcho2 (0.00s)
	=== RUN   TestEcho3
	/tmp/go-build1718110061/b001/ch1ex1.test -test.paniconexit0 -test.timeout=10m0s -test.v=true
	--- PASS: TestEcho3 (0.00s)
	PASS
	ok      github.com/aspieunixgeek/go-exercises/ch1ex1    0.002s

*/

package ch1ex1
