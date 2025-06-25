// Exercise 2.4:
// Write a version of PopCount that counts bits by shifting its argument through 64 bit positions,
// testing the rightmost bit each time.
// Compare its performance to the table-lookup version.

/*
# Run unit tests

	$ go test -v
	=== RUN   Test_PopCount
	    ch2ex4_test.go:13: x: 4875148618339312146
	--- PASS: Test_PopCount (0.00s)
	PASS
	ok      github.com/aspieunixgeek/go-exercises/ch2ex4    0.002s

# Benchmark_PopCountOld (0.366s)

	$ go test -v -bench="Old" | tail
		ch2ex4_test.go:13: x: 17762119341007886997
	--- PASS: Test_PopCount (0.00s)
	goos: linux
	goarch: amd64
	pkg: github.com/aspieunixgeek/go-exercises/ch2ex4
	cpu: 12th Gen Intel(R) Core(TM) i5-12450H
	Benchmark_PopCountOld
	Benchmark_PopCountOld-12        1000000000               0.3264 ns/op
	PASS
	ok      github.com/aspieunixgeek/go-exercises/ch2ex4    0.366s

# Benchmark_PopCountNew (1.488s)

	$ go test -v -bench="New" | tail
		ch2ex4_test.go:13: x: 7241454065380347684
	--- PASS: Test_PopCount (0.00s)
	goos: linux
	goarch: amd64
	pkg: github.com/aspieunixgeek/go-exercises/ch2ex4
	cpu: 12th Gen Intel(R) Core(TM) i5-12450H
	Benchmark_PopCountNew
	Benchmark_PopCountNew-12        491705575                2.510 ns/op
	PASS
	ok      github.com/aspieunixgeek/go-exercises/ch2ex4    1.488s
*/
package ch2ex4
