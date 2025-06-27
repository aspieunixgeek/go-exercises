// Exercise 2.5:
// The expression x&(x-1) clears the rightmost non-zero bit of x.
// Write a version of PopCount that counts bits by using this fact,
// and assess its performance.
/*

# PopCountOld

	$ go test -bench="Old"
	goos: linux
	goarch: amd64
	pkg: github.com/aspieunixgeek/go-exercises/ch2ex5
	cpu: 12th Gen Intel(R) Core(TM) i5-12450H
	BenchmarkOld-12         1000000000               0.2359 ns/op
	PASS
	ok      github.com/aspieunixgeek/go-exercises/ch2ex5    0.270s

# PopCountNew

	$ go test -bench="New"
	goos: linux
	goarch: amd64
	pkg: github.com/aspieunixgeek/go-exercises/ch2ex5
	cpu: 12th Gen Intel(R) Core(TM) i5-12450H
	BenchmarkNew-12         46618694                28.92 ns/op
	PASS
	ok      github.com/aspieunixgeek/go-exercises/ch2ex5    1.379s


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
