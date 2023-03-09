package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Square struct {
	width  float64
	height float64
}

func (s Square) area() float64 {
	return s.width * s.height
}

func main() {
	c1 := Circle{radius: 4}
	s1 := Square{width: 10, height: 10}

	shapes := []Shape{c1, s1}
	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
}
