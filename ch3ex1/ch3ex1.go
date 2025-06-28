package ch3ex1

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

func Gen() string {
	out := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok := corner(i+1, j)
			if !ok {
				break
			}
			bx, by, ok := corner(i, j)
			if !ok {
				break
			}
			cx, cy, ok := corner(i, j+1)
			if !ok {
				break
			}

			dx, dy, ok := corner(i+1, j+1)
			if !ok {
				break
			}
			out += fmt.Sprintf("\n<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	out += fmt.Sprintf("\n</svg>")

	return out
}

func corner(i, j int) (x float64, y float64, ok bool) {
	// Find point (x,y) at corner of cell (i,j).
	x = xyrange * (float64(i)/cells - 0.5)
	y = xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, ok := f(x, y)
	if !ok {
		return -1, -1, false
	}
	// Project (x, y, z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) (z float64, ok bool) {
	r := math.Hypot(x, y)  // distance from (0,0)
	if !math.IsInf(r, 0) { // If sign == 0, IsInf reports whether f is either infinity.
		if math.IsInf(math.Sin(r)/r, 0) {
			return // A "return" statement in a function F terminates the execution of F
		} else {
			return math.Sin(r) / r, true
		}
	}
	return // A "return" statement in a function F terminates the execution of F
}
