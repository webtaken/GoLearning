package main

import "fmt"

func main() {
	a := 11
	b := 500

	fmt.Printf("mul1(%[1]d,%[2]d): %d\nmul2(%[1]d,%[2]d): %d\n", a, b, Mul1(a, b), Mul2(a, b))
}

func Mul1(a, b int) int {
	ans := 0
	pivot_a := a
	for b != 0 {
		if (b & 1) == 1 {
			ans += pivot_a
		}
		b = b >> 1
		pivot_a += pivot_a
	}
	return ans
}

func Mul2(a, b int) int {
	ans := 0
	for i := 0; i < b; i++ {
		ans += a
	}
	return ans
}
