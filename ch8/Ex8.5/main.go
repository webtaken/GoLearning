// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"sync"
	"time"
)

var mapping [16]color.RGBA

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No provided name of the PNG file")
	}

	mapping[0] = color.RGBA{66, 30, 15, 255}
	mapping[1] = color.RGBA{25, 7, 26, 255}
	mapping[2] = color.RGBA{9, 1, 47, 255}
	mapping[3] = color.RGBA{4, 4, 73, 255}
	mapping[4] = color.RGBA{0, 7, 100, 255}
	mapping[5] = color.RGBA{12, 44, 138, 255}
	mapping[6] = color.RGBA{24, 82, 177, 255}
	mapping[7] = color.RGBA{57, 125, 209, 255}
	mapping[8] = color.RGBA{134, 181, 229, 255}
	mapping[9] = color.RGBA{211, 236, 248, 255}
	mapping[10] = color.RGBA{241, 233, 191, 255}
	mapping[11] = color.RGBA{248, 201, 95, 255}
	mapping[12] = color.RGBA{255, 170, 0, 255}
	mapping[13] = color.RGBA{204, 128, 0, 255}
	mapping[14] = color.RGBA{153, 87, 0, 255}
	mapping[15] = color.RGBA{106, 52, 3, 255}

	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	f, _ := os.Create(os.Args[1])

	start := time.Now()
	var wg sync.WaitGroup
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		wg.Add(1)
		go func(py int, y float64) {
			defer wg.Done()
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				// Image point (px, py) represents complex value z.
				img.Set(px, py, mandelbrot(z))
			}
		}(py, y)
	}

	wg.Wait()

	// Now we will computer supersampling
	for py := 0; py < height; py += 2 {
		wg.Add(1)
		go func(py int) {
			defer wg.Done()
			for px := 0; px < width; px += 2 {
				r1, g1, b1, _ := img.At(px, py).RGBA()
				r2, g2, b2, _ := img.At(px+1, py).RGBA()
				r3, g3, b3, _ := img.At(px, py+1).RGBA()
				r4, g4, b4, _ := img.At(px+1, py+1).RGBA()

				avgRed := (r1 + r2 + r3 + r4) / 4
				avgGreen := (g1 + g2 + g3 + g4) / 4
				avgBlue := (b1 + b2 + b3 + b4) / 4

				pxAvg := color.RGBA{uint8(avgRed), uint8(avgGreen), uint8(avgBlue), 255}

				img.Set(px, py, pxAvg)
				img.Set(px+1, py, pxAvg)
				img.Set(px, py+1, pxAvg)
				img.Set(px+1, py+1, pxAvg)
			}
		}(py)
	}
	wg.Wait()
	png.Encode(f, img) // NOTE: ignoring errors
	timeElapsed := time.Since(start)
	fmt.Printf("Execution time took %d(ms)", timeElapsed.Milliseconds())
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		modulus := cmplx.Abs(v)
		if modulus > 2 {
			iter := n % 16
			return mapping[iter]
		}
	}
	return color.Black
}
