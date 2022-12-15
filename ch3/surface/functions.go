package main

import "math"

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	if math.IsNaN(math.Sin(r) / r) {
		return 0, false
	}
	return math.Sin(r) / r, true
}

func eggbox(x, y float64) (float64, bool) {
	if math.IsNaN(0.2 * (math.Cos(x) + math.Cos(y))) {
		return 0, false
	}
	return 0.2 * (math.Cos(x) + math.Cos(y)), true
}

func saddle(x, y float64) (float64, bool) {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	if math.IsNaN(y*y/a2 - x*x/b2) {
		return 0, false
	}
	return (y*y/a2 - x*x/b2), true
}
