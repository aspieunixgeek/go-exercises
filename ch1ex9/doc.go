// Exercise 1.9:
// Modify fetch to also print the HTTP status code, found in resp.Status.

/*
# Go mod init

	$ go mod init ch1ex9

# Unit test output

	$ go test -v
	=== RUN   TestFetch

	ch1ex9_test.go:21: <!-- 200 OK -->

	--- PASS: TestFetch (1.36s)
	PASS
	ok      ch1ex9  1.369s
*/
package ch1ex9
