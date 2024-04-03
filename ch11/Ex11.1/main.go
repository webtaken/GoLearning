package main

import (
	"fmt"
	"strings"

	"Ex11.1/charcount"
)

func main() {
	newRd := strings.NewReader("abcdefg")
	counts, utflen, invalid := charcount.CharCount(newRd)
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
