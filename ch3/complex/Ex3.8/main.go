// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	fractalColors "fractals/FractalColors"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/big"
	"math/cmplx"
	"os"
)

type ComplexF struct {
	r big.Float
	i big.Float
}

func MulComplexF(a, b *ComplexF) ComplexF {
	real := a.r.Sub(a.r.Mul(&a.r, &a.r), a.i.Mul(&a.i, &a.i))
	imag := a.r.Mul(a.r.Mul(&a.r, big.NewFloat(2)), &a.i)
	return ComplexF{*real, *imag}
}

func DivComplexF(a, b *ComplexF) ComplexF {
	denominator := b.r.Add(b.r.Mul(&b.r, &b.r), b.i.Mul(&b.i, &b.i))
	real := a.r.Quo(a.r.Add(a.r.Mul(&a.r, &b.r), b.i.Mul(&a.i, &b.i)), denominator)
	imag := a.r.Quo(a.r.Sub(b.r.Mul(&b.r, &a.i), a.r.Mul(&a.r, &b.i)), denominator)
	return ComplexF{*real, *imag}
}

func AddComplexF(a, b *ComplexF) ComplexF {
	real := a.i.Add(&a.r, &b.r)
	imag := a.i.Add(&a.i, &b.i)
	return ComplexF{*real, *imag}
}

func SubComplexF(a, b *ComplexF) ComplexF {
	real := a.r.Sub(&a.r, &b.r)
	imag := a.i.Sub(&a.i, &b.i)
	return ComplexF{*real, *imag}
}

func ModulusComplexF(a ComplexF) big.Float {
	var tmpr, tmpi, sum, sqrt *big.Float
	tmpr = tmpr.Mul(&a.r, &a.r)
	tmpi = tmpi.Mul(&a.i, &a.i)
	sum = sum.Add(tmpr, tmpi)
	sqrt = sqrt.Sqrt(sum)
	return *sqrt
}

func PrintComplexF(a ComplexF) {
	real := a.r.Add(&a.r, big.NewFloat(0))
	imag := a.i.Add(&a.i, big.NewFloat(0))
	fmt.Printf("%v + %vi", real, imag)
}

func main() {
	var a = ComplexF{*big.NewFloat(3), *big.NewFloat(5)}
	var b = ComplexF{*big.NewFloat(5), *big.NewFloat(9)}
	var c = AddComplexF(&a, &b)
	PrintComplexF(a)
	fmt.Println()
	PrintComplexF(b)
	fmt.Println()
	PrintComplexF(c)
	fmt.Println()
	return
	// usage ./bin <mandelbrot|newton> <64|128|F|R> <image.png>

	if len(os.Args) < 4 {
		log.Fatal("No provided enough arguments, usage: ./bin <mandelbrot|newton> <64|128|F|R> <image.png>")
	}

	fractalType := os.Args[1]

	if fractalType != "newton" && fractalType != "mandelbrot" {
		log.Fatal("There are only two types of fractals: <mandelbrot|newton>.")
	}

	precisionType := os.Args[2]

	if precisionType != "64" && precisionType != "128" && precisionType != "F" && precisionType != "R" {
		log.Fatal("There are only two types of fractals: <64|128|F|R>.")
	}

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
			switch precisionType {
			case "64":
				if fractalType == "newton" {
					img.Set(px, py, z464(complex64(z)))
				} else {
					img.Set(px, py, mandelbrot64(complex64(z)))
				}
			case "128":
				if fractalType == "newton" {
					img.Set(px, py, z4128(z))
				} else {
					img.Set(px, py, mandelbrot128(z))
				}
			case "F":
				var zF = ComplexF{*big.NewFloat(x), *big.NewFloat(y)}
				if fractalType == "newton" {
					img.Set(px, py, z4BigF(zF))
				} else {
					img.Set(px, py, mandelbrotBigF(zF))
				}

				img.Set(px, py, z4128(z))
			case "R":
				img.Set(px, py, z4128(z))
			}
		}
	}

	png.Encode(f, img) // NOTE: ignoring errors
}

func mandelbrot128(z complex128) color.Color {
	const iterations = 200
	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		modulus := cmplx.Abs(v)
		if modulus > 2 {
			iter := n % 16
			return fractalColors.Mapping[iter]
		}
	}
	return color.Black
}

func mandelbrot64(z complex64) color.Color {
	const iterations = 200
	var v complex64

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		modulus := cmplx.Abs(complex128(v))
		if modulus > 2 {
			iter := n % 16
			return fractalColors.Mapping[iter]
		}
	}
	return color.Black
}

func mandelbrotBigF(z ComplexF) color.Color {

	const iterations = 200
	var v = ComplexF{*big.NewFloat(0), *big.NewFloat(0)}
	for n := uint8(0); n < iterations; n++ {
		vPow2 := MulComplexF(&v, &v)
		v = AddComplexF(&vPow2, &z)
		mod := ModulusComplexF(v)
		if mod.Cmp(big.NewFloat(2)) < 0 { // this means mod < 2
			iter := n % 16
			return fractalColors.Mapping[iter]
		}
	}
	return color.Black
}

func z4128(z complex128) color.Color {
	const iterations = 200
	var v complex128 = z

	for n := uint8(0); n < iterations; n++ {
		v = newton128(v, v)
		val := v*v*v*v - 1
		if cmplx.Abs(val) < 1e-3 {
			iter := n % 16
			return fractalColors.Mapping[iter]
		}
	}
	return color.Black
}

func z464(z complex64) color.Color {
	const iterations = 200
	var v complex64 = z

	for n := uint8(0); n < iterations; n++ {
		v = newton64(v, v)
		val := v*v*v*v - 1
		if cmplx.Abs(complex128(val)) < 1e-3 {
			iter := n % 16
			return fractalColors.Mapping[iter]
		}
	}
	return color.Black
}

func z4BigF(z ComplexF) color.Color {
	const iterations = 200
	var v = ComplexF{*big.NewFloat(0), *big.NewFloat(0)}
	var one = ComplexF{*big.NewFloat(1), *big.NewFloat(0)}

	for n := uint8(0); n < iterations; n++ {
		v = newtonBigF(v, v)
		var zPow2 = MulComplexF(&v, &v)              // z^2
		var zPow3 ComplexF = MulComplexF(&v, &zPow2) // z^3
		var zPow4 ComplexF = MulComplexF(&v, &zPow3) // z^4

		val := SubComplexF(&zPow4, &one)
		mod := ModulusComplexF(val)
		if mod.Cmp(big.NewFloat(1e-3)) < 0 { // mod < 1e-3
			iter := n % 16
			return fractalColors.Mapping[iter]
		}
	}
	return color.Black
}

// P(z) = z^4 - 1
// P'(z) = 4*z^3
// x_(n+1) = x_n - P(z) / P'(z) = x_n - 1/4(z - (1/z^3))
func newton128(x_n, z complex128) complex128 {
	return x_n - 1.0/4*(z-(1/(z*z*z)))
}

func newton64(x_n, z complex64) complex64 {
	return x_n - 1.0/4*(z-(1/(z*z*z)))
}

func newtonBigF(x_n, z ComplexF) ComplexF {
	var a = ComplexF{*big.NewFloat(0.25), *big.NewFloat(0)}
	var one = ComplexF{*big.NewFloat(1), *big.NewFloat(0)}
	var quadraticZ = MulComplexF(&z, &z)               // z^2
	var cubicZ ComplexF = MulComplexF(&z, &quadraticZ) // z^3
	var divisionZ = DivComplexF(&one, &cubicZ)         // 1/z^3
	var subZ ComplexF = SubComplexF(&z, &divisionZ)    // z - (1/z^3)
	deriv := MulComplexF(&a, &subZ)                    // 1/4 * (z-(1/z^3))
	return SubComplexF(&x_n, &deriv)
}
