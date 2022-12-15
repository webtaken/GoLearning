// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"image/color"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

// var palette = []color.Color{
// 	color.White,
// 	color.Black,
// 	color.RGBA{0x00, 0xff, 0x00, 0xff},
// 	color.RGBA{0x00, 0xff, 0xff, 0xff},
// 	color.RGBA{0x01, 0x08, 0xaa, 0xff}}

var (
	zMin, zMax    = math.Inf(1), math.Inf(-1) // minimun and maximun
	zColor        = color.RGBA{0, 0, 0, 1}
	width, height = 600, 320
	// canvas size in pixels
	cells = 100
	// number of grid cells
	xyrange = 30.0
	// axis ranges (-xyrange..+xyrange)
	xyscale = float64(width) / 2 / xyrange // pixels per x or y unit
	zscale  = float64(height) * 0.4
	// pixels per z unit
	angle    = math.Pi / 6 // angle of x, y axes (=30°)
	funcType = "f"         // type of function "f" or "eggbox" or "saddle"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		qWidth, err := strconv.Atoi(q["w"][0])
		if err != nil {
			fmt.Fprintf(w, "Error: %s\n", err)
		}
		width = qWidth
		qHeight, err := strconv.Atoi(q["h"][0])
		if err != nil {
			fmt.Fprintf(w, "Error: %s\n", err)
		}
		height = qHeight
		fType := q["t"][0]
		funcType = fType
		w.Header().Set("Content-Type", "image/svg+xml")
		svgGenerator(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)
	// Compute surface height z.
	z, ok := float64(0.0), true
	switch funcType {
	case "f":
		z, ok = f(x, y)
	case "eggbox":
		z, ok = eggbox(x, y)
	case "saddle":
		z, ok = saddle(x, y)
	default:
		z, ok = f(x, y)
	}

	t := normalize(z, zMin, zMax)
	zColor.R = uint8(t * 255)
	zColor.B = uint8((1 - t) * 255)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, ok
}

func getBounds() (float64, float64) {
	minima, maxima := math.Inf(1), math.Inf(-1)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			// Find point (x,y) at corner of cell (i,j).
			x := xyrange * (float64(i)/float64(cells) - 0.5)
			y := xyrange * (float64(j)/float64(cells) - 0.5)
			// Compute surface height z.
			z, ok := float64(0.0), true
			switch funcType {
			case "f":
				z, ok = f(x, y)
			case "eggbox":
				z, ok = eggbox(x, y)
			case "saddle":
				z, ok = saddle(x, y)
			default:
				z, ok = f(x, y)
			}

			if ok {
				if z < minima {
					minima = z
				}
				if z > maxima {
					maxima = z
				}
			}
		}
	}
	return minima, maxima
}

// Normalizes z between [0,1], where 1 is the max and 0 is the minimun
func normalize(z, zmin, zmax float64) float64 {
	return (z - zmin) / (zmax - zmin)
}

func svgGenerator(out io.Writer) {
	svgImage := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	zMin, zMax = getBounds()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok1 := corner(i+1, j)
			bx, by, ok2 := corner(i, j)
			cx, cy, ok3 := corner(i, j+1)
			dx, dy, ok4 := corner(i+1, j+1)
			if !(ok1 && ok2 && ok3 && ok4) {
				continue
			}
			selectedColor := fmt.Sprintf("#%x%x%x", zColor.R, zColor.G, zColor.B)
			svgImage += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: %s; fill:#222222'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, selectedColor)
		}
	}
	svgImage += "</svg>"
	fmt.Printf("%g %g\n", zMin, zMax)
	fmt.Fprint(out, svgImage)
}
