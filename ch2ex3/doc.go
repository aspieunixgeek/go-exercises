// Exercise 2.3:
// Rewrite PopCount to use a loop instead of a single expression.
// Compare the performance of the two versions.
// (Section 11.4 shows how to compare the performance of different implementations systematically.)

/*
# Run benchmarks

	$ go test -bench="PopCount"
	goos: linux
	goarch: amd64
	pkg: github.com/aspieunixgeek/go-exercises/ch2ex3
	cpu: 12th Gen Intel(R) Core(TM) i5-12450H
	BenchmarkPopCountOld-12         100000000               10.81 ns/op
	BenchmarkPopCountNew-12         70095602                16.97 ns/op
	PASS
	ok      github.com/aspieunixgeek/go-exercises/ch2ex3    2.306s

# Run unit tests

	$ go test -v
	=== RUN   TestPopCount
	--- PASS: TestPopCount (0.00s)
	=== RUN   TestPopCount2
	--- PASS: TestPopCount2 (0.00s)
	PASS
	ok      github.com/aspieunixgeek/go-exercises/ch2ex3    0.002s
*/
package ch2ex3
