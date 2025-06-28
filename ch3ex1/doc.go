// Exercise 3.1:
// If the function f returns a non-finite float64 value,
// the SVG file will contain invalid <polygon> elements (although many SVG renderers handle this gracefully).
// Modify the program to skip invalid polygons.

/*
# Unit Test

	$ go test -v
	=== RUN   Test_ChDir
	--- PASS: Test_ChDir (0.00s)
	=== RUN   Test_Gen
	--- PASS: Test_Gen (2.30s)
	PASS
	ok      github.com/aspieunixgeek/go-exercises/ch3ex1    2.308s

# Change dir to out

	$ cd out

# Run ls

	$ ls
	out_28_06_2025__17_28_56.svg
*/
package ch3ex1
