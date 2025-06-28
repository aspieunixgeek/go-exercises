package ch3ex2

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func Gen(f func(x, y float64) float64) string {
	var out string

	out = fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			out += fmt.Sprintf("\n<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	out += fmt.Sprintf("\n</svg>")

	return out
}

func corner(i, j int, f func(x, y float64) float64) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y) //f1(x, y)

	// Project (x, y, z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f1(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func f2(x, y float64) float64 {
	r := math.Dim(x, y) // Dim returns the maximum of x-y or 0.
	return math.Sin(r) / r
}

func f3(x, y float64) float64 {
	r := math.Mod(x, y) // Mod returns the floating-point remainder of x/y. The magnitude of the result is less than y and its sign agrees with that of x.
	return math.Sin(r) / r
}

func f4(x, y float64) float64 {
	r := math.Max(x, y) // Max returns the larger of x or y.
	return math.Sin(r) / r
}

func f5(x, y float64) float64 {
	r := math.Min(x, y) // Min returns the smaller of x or y.
	return math.Sin(r) / r
}

func f6(x, y float64) float64 {
	r := math.Nextafter(x, y) // Nextafter returns the next representable float64 value after x
	return math.Sin(r) / r
}
