// Cf converts its numeric argument to Celsius and Fahrenheit.

/*
# Init go-module

	$ go mod init cf
	go: creating new go.mod: module cf
	go: to add module requirements and sums:
		go mod tidy

# Install dependencies

	$ go get github.com/aspieunixgeek/go-exercises/tempconv
	go: downloading github.com/aspieunixgeek/go-exercises v0.0.0-20250618192608-f3ef120a3028
	go: added github.com/aspieunixgeek/go-exercises v0.0.0-20250618192608-f3ef120a3028

# Run

	$ go run ./main.go 8 16 32
	8°F = -13.333333333333334°C, 8°C = 46.4°F
	16°F = -8.88888888888889°C, 16°C = 60.8°F
	32°F = 0°C, 32°C = 89.6°F
*/
package main
