// Exercise 1.3:
// Experiment to measure the difference in running time between our potentially inefficient versions
// and the one that uses strings.Join. (Section 1.6 illustrates part of the time package,
// and Section 11.4 shows how to write benchmark tests for systematic performance evaluation.)

/*

# Echo1

	$ go test -bench="Echo1" | tail
	-test.paniconexit0 -test.timeout=10m0s -test.bench=Echo1
	-test.paniconexit0 -test.timeout=10m0s -test.bench=Echo1
	-test.paniconexit0 -test.timeout=10m0s -test.bench=Echo1
	-test.paniconexit0 -test.timeout=10m0s -test.bench=Echo1
	-test.paniconexit0 -test.timeout=10m0s -test.bench=Echo1
	-test.paniconexit0 -test.timeout=10m0s -test.bench=Echo1
	-test.paniconexit0 -test.timeout=10m0s -test.bench=Echo1
	 1000000              1034 ns/op
	PASS
	ok      github.com/aspieunixgeek/go-exercises/ch1ex3    1.050s


# Echo2

	$ go test -bench="Echo2" | tail
	-test.paniconexit0 -test.timeout=10m0s -test.bench=Echo2
	-test.paniconexit0 -test.timeout=10m0s -test.bench=Echo2
	-test.paniconexit0 -test.timeout=10m0s -test.bench=Echo2
	-test.paniconexit0 -test.timeout=10m0s -test.bench=Echo2
	-test.paniconexit0 -test.timeout=10m0s -test.bench=Echo2
	-test.paniconexit0 -test.timeout=10m0s -test.bench=Echo2
	-test.paniconexit0 -test.timeout=10m0s -test.bench=Echo2
	 1204549               982.3 ns/op
	PASS
	ok      github.com/aspieunixgeek/go-exercises/ch1ex3    2.195s


# Echo3

	$ go test -bench="Echo3" | tail
	-test.paniconexit0 -test.timeout=10m0s -test.bench=Echo3
	-test.paniconexit0 -test.timeout=10m0s -test.bench=Echo3
	-test.paniconexit0 -test.timeout=10m0s -test.bench=Echo3
	-test.paniconexit0 -test.timeout=10m0s -test.bench=Echo3
	-test.paniconexit0 -test.timeout=10m0s -test.bench=Echo3
	-test.paniconexit0 -test.timeout=10m0s -test.bench=Echo3
	-test.paniconexit0 -test.timeout=10m0s -test.bench=Echo3
	 1349220               901.1 ns/op
	PASS
	ok      github.com/aspieunixgeek/go-exercises/ch1ex3    2.120s

*/

package ch1ex3
