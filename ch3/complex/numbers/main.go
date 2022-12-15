package main

import (
	"fmt"
)

func main() {
	x := 1 + 2.5i
	y := 1.3i - 5
	fmt.Printf("%v + %v = %v\n", x, y, x+y)
	fmt.Printf("%v - %v = %v\n", x, y, x-y)
	fmt.Printf("%v * %v = %v\n", x, y, x*y)
	fmt.Printf("%v / %v = %v\n", x, y, x/y)
}
