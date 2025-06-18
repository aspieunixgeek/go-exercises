// Exercise 2.2:
// Write a general-purpose unit-conversion program analogous to cf that
// reads numbers from its command-line arguments or from the standard input if there are no arguments,
// and converts each number into like temperature in Celsius and Fahrenheit, length in feet and meters,
// weight in pounds and kilograms, and the like.

/*
# Output

	$ go run ./ch2ex2.go 8 16 32
	----------------------------------------------------------------
	8°F = -13.333333333333334°C, 8°C = 46.4°F
	8.00 ft. = 2.44 M, 8.00 M = 26.25 ft.
	8.00 Kg = 17.64 lb., 8.00 lb. = 3.63 Kg
	8.00 ps = 0.01 ns, 8.00 ns = 8000.00 ps
	----------------------------------------------------------------
	16°F = -8.88888888888889°C, 16°C = 60.8°F
	16.00 ft. = 4.88 M, 16.00 M = 52.50 ft.
	16.00 Kg = 35.28 lb., 16.00 lb. = 7.26 Kg
	16.00 ps = 0.02 ns, 16.00 ns = 16000.00 ps
	----------------------------------------------------------------
	32°F = 0°C, 32°C = 89.6°F
	32.00 ft. = 9.75 M, 32.00 M = 104.99 ft.
	32.00 Kg = 70.56 lb., 32.00 lb. = 14.51 Kg
	32.00 ps = 0.03 ns, 32.00 ns = 32000.00 ps
	----------------------------------------------------------------
*/
package main
