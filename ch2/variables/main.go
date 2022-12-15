package main

import (
	"fmt"
)

func main() {
	var integer, floating, str = 5, 3.4, "hola mundo" // multiple declarations in one line
	fmt.Printf("%v %v %v\n", integer, floating, str)

	a, b := 4, 5
	fmt.Printf("%v\n", a+b)
	a, c := 7, 8
	fmt.Printf("%v\n", a+c)

	x := 5
	p := &x
	p2 := &p
	fmt.Printf("Type of p: %T\n", p)
	fmt.Printf("Type of p2: %T\n", p2)
	fmt.Printf("(%v)(%v)\t(%v)(%v)\n", *p2, p2, p, &p)
	fmt.Printf("Value of p(%v) and the address %v\n", *p, p)

	my_p1 := f()
	my_p2 := f()
	fmt.Printf("%v %v\n", my_p1, my_p2)
	fmt.Printf("%v\n", my_p1 == my_p2)
}

func f() *int {
	v := 0
	return &v
}
