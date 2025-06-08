// Exercise 1.8:
// Modify fetch to add the prefix http:// to each argument URL if it is missing. You might want to use strings.HasPrefix.

/*
# Init go mod file

	$ go mod init ch1ex8

# Unit tests

	$ go test -v
	=== RUN   TestFetch1
	--- PASS: TestFetch1 (2.15s)
	=== RUN   TestFetch2
	--- PASS: TestFetch2 (0.75s)
	=== RUN   TestFetch3
	--- PASS: TestFetch3 (0.90s)
	PASS
	ok      ch1ex8  3.794s
*/
package ch1ex8
