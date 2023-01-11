package main

import (
	"fmt"
	"unicode/utf8"
)

func rev(a []byte) {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-1-i] = a[len(a)-1-i], a[i]
	}
}

func reverse(a []byte) []byte {
	for i := 0; i < len(a); {
		_, size := utf8.DecodeRune(a[i:])
		rev(a[i : i+size])
		i += size
	}
	rev(a)
	return a
}

func main() {
	msg := "hola como estas"
	encodedUTF8 := []byte(msg)
	fmt.Printf("Message: %s\nEncoded: %v\n", msg, encodedUTF8)
	reverse(encodedUTF8)
	fmt.Printf("Reversed: %v\n", encodedUTF8)
}
