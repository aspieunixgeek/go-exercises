// Exercise 1.2:
// Modify the echo program to print the index and value of each of its arguments, one per line.

/*
# Run without flags

	$ go test
	index: 0, value: a
	index: 1, value: b
	index: 2, value: c

	index: 0, value: a
	index: 1, value: b
	index: 2, value: c

	index: 0, value: a
	index: 1, value: b
	index: 2, value: c
	PASS
	ok      github.com/aspieunixgeek/go-exercises/ch1ex2    0.002s

# Run with v flag

	$ go test -v
	=== RUN   TestEcho1
	index: 0, value: a
	index: 1, value: b
	index: 2, value: c

	--- PASS: TestEcho1 (0.00s)
	=== RUN   TestEcho2
	index: 0, value: a
	index: 1, value: b
	index: 2, value: c

	--- PASS: TestEcho2 (0.00s)
	=== RUN   TestEcho3
	index: 0, value: a
	index: 1, value: b
	index: 2, value: c
	--- PASS: TestEcho3 (0.00s)
	PASS
	ok      github.com/aspieunixgeek/go-exercises/ch1ex2    0.003s

*/

package ch1ex2
