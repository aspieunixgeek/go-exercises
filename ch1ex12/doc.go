// Exercise 1.12:
// Modify the Lissajous server to read parameter values from the URL.
// For example, you might arrange it so that a URL like http://localhost:8000/?cycles=20 sets
// the number of cycles to 20 instead of the default 5. Use the strconv.Atoi function to
// convert the string parameter into an integer. You can see its documentation with go doc
// strconv.Atoi.

/*
# Change current dir

	cd ch1ex12

# Init go module

	$ go mod init ch1ex12

# Run

	$ go run ./main.go &

# Open web-browser

	localhost:8000
*/
package main
