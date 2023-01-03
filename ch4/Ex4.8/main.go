// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	category := map[string]uint{
		"Letter (L)":      0,
		"Mark (M)":        0,
		"Number (N)":      0,
		"Punctuation (P)": 0,
		"Symbol (S)":      0,
		"Separator (Z)":   0,
		"Other (O)":       0,
	}
	// counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0
	// count of invalid UTF-8 characters
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		if unicode.IsLetter(r) {
			category["Letter (L)"]++
		} else if unicode.IsMark(r) {
			category["Mark (M)"]++
		} else if unicode.IsNumber(r) {
			category["Number (N)"]++
		} else if unicode.IsPunct(r) {
			category["Punctuation (P)"]++
		} else if unicode.IsSymbol(r) {
			category["Symbol (S)"]++
		} else if unicode.IsSpace(r) {
			category["Separator (Z)"]++
		} else {
			category["Other (O)"]++
		}
	}
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
	fmt.Printf("\nunicode category\n")
	for c, n := range category {
		fmt.Printf("%q: %d found\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
