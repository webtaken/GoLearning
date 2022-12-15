// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
)

var mapping [16]color.RGBA

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No provided name of the image file, usage: ./bin <FILENAME>")
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
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, z4(z))
		}
	}

	png.Encode(f, img) // NOTE: ignoring errors
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

func z4(z complex128) color.Color {
	const iterations = 200
	var v complex128 = z

	for n := uint8(0); n < iterations; n++ {
		v = newton(v, v)
		val := v*v*v*v - 1
		if cmplx.Abs(val) < 1e-3 {
			iter := n % 16
			return mapping[iter]
		}
	}
	return color.Black
}

// P(z) = z^4 - 1
// P'(z) = 4*z^3
// x_(n+1) = x_n - P(z) / P'(z) = x_n - 1/4(z - (1/z^3))
func newton(x_n, z complex128) complex128 {
	return x_n - 1.0/4*(z-(1/(z*z*z)))
}
