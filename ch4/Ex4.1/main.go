package main

import (
	"crypto/sha256"
	"fmt"

	"ex4.1/compare"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%[1]v\n%[2]v\n%[1]x\n%[2]x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	a := byte(5)  // 00000101
	b := byte(15) // 00001111
	fmt.Printf("%v\n", compare.Compare2Bytes(a, b))
	fmt.Printf("N° different bytes: %v\n", compare.CompareSHA256(c1, c2))
}
