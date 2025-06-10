// Exercise 1.10:
// Find a web site that produces a large amount of data. Investigate caching by running fetchall twice in
// succession to see whether the reported time changes much. Do you get the same content each time?
// Modify fetchall to print its output to a file so it can be examined.

/*
# Go mod init

	$ go mod init ch1ex10

# Clear cache

	go clean -testcache

# Run tests twice

	$ go test
	PASS
	ok      ch1ex10 2.762s

	$ go test
	PASS
	ok      ch1ex10 2.634s

# Help

	$ go help cache
*/
package ch1ex10
