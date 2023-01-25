package main

import "fmt"

// squares returns a function that returns
// the next square number each time it is called.
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func main() {
	f1 := squares()
	f2 := squares()
	for i := 0; i < 10; i++ {
		fmt.Printf("f1: %v\tf2: %v\n", f1(), f2())
	}
}
