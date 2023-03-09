package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	p1 := Point{0, 0}
	p2 := Point{3, 4}
	p3 := Point{4, 4}
	path := Path{p1, p2, p3}

	fmt.Printf("%v\n", path.Distance())
	fmt.Printf("%v\n", Distance(p1, p2))
	fmt.Printf("%v\n", p1.Distance(p2))

}
