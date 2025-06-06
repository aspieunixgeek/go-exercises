package ch1ex6

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF}, // black background color
	color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}, // red foreground color
	color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF}, // green foreground color
	color.RGBA{R: 0x00, G: 0x00, B: 0xFF, A: 0xFF}, // blue foreground color
}

const (
	blackIndex = 0 // background color in palette
	redIndex   = 1 // first color in palette
	greenIndex = 2 // third color in palette
	blueIndex  = 3 // four color in palette
)

var curInd uint8 = 1

func Gen() error {

	for i := 1; i < 4; i++ {
		filename := fmt.Sprintf("file_%d.gif", i)

		f, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("os.Create: %v\n", err)
		}

		lissajous(f)
		curInd++
	}

	return nil
}
func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complette x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), curInd)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
