package main

import (
	"fmt"
	"image/color"
)

type Dim struct{ dim int }
type Point struct{ X, Y float64 }
type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p *Point) ScaleBy(x float64) {
	p.X *= x
	p.Y *= x
}

func main() {
	p1 := Point{1, 2}
	(&p1).ScaleBy(2)
	fmt.Printf("%v", p1)

	var cp ColoredPoint
	cp

}
