// Exercise 3.4:
// Following the approach of the Lissajous example in Section 1.7, construct a web server that computes surfaces and
// writes SVG data to the client. The server must set the Content-Type header like this:
/*
   w.Header().Set("Content-Type", "image/svg+xml")
*/
// (This step was not required in the  Lissajous example because the server uses standard heuristics
// to recognise common formats like PNG from the first 512 bytes of the response, and generates the proper header.)
// Allow the client to specify values like height, width, and color as HTTP request parameters.

/*
# Build

	$ go build

# Usage with default port

	$ ./ch3ex4

# Usage with custom port

	$ ./ch3ex4 -p 9000

#  Browser

	http://localhost:9000/?color=blue&bg=black&width=1024&height=1024
*/
package main
