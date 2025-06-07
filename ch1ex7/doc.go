// Exercise 1.7:
// The function call io.Copy(dst, src) reads from src and writes to dst.
// Use it instead of io.ReadAll to copy the response body to os.Stdout without requiring a buffer large enough to hold the
// entire stream. Be sure to check the error result of io.Copy.

/*
# Init go module

	$ go mod init ch1ex7

# Run unit test

	$ go test
	...stdout output
	PASS
	ok      ch1ex7  2.433s

*/

package ch1ex7
