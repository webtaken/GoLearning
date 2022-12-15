package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	var a int = 32
	var ru rune = 21
	fmt.Printf("%v\n", unsafe.Sizeof(a))
	fmt.Printf("%v\n", unsafe.Sizeof(ru))
	fmt.Println(5.0 / 4)
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // "255 0 1"
	var i int8 = 127
	fmt.Println(i, i+1, i*i) // "127 -128 1"
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
	// "00100010", the set {1, 5}
	// "00000110", the set {1, 2}
	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}
	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}

	f := 1e100  // a float64
	n := int(f) // result is implementation-dependent
	fmt.Println(n)

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	x1 := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x1)

	ascii := 'a'
	unicode := '🤣'
	newline := '\n'
	fmt.Printf("%d %[1]T %[1]c %[1]q\n", ascii)
	// "97 a 'a'"
	fmt.Printf("%d %[1]T %[1]c %[1]q\n", unicode) // "22269 D 'D'"
	fmt.Printf("%d %[1]T %[1]q\n", newline)
	// "10 '\n'"

	var f1 float32 = 16777216 // 1 << 24
	fmt.Println(f1 == f1+1)
	// "true"!

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d ℯ = %8.3f\n", x, math.Exp(float64(x)))
	}
}
