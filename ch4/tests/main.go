package main

import "fmt"

func main() {
	a := [3]byte{4, 2, 240}
	fmt.Printf("%T\t%[1]v\n", a)
	var ptr *[3]byte = &a
	zeroes(ptr)
	fmt.Printf("%T\t%[1]v\n", a)

	strArr := []string{"one", "", "three"}

	fmt.Printf("%v\n", strArr[0:0])
}

func zeroes(ptr *[3]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}
