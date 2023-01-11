package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squashSpaces(s []byte) []byte {
	i := 0
	ans := []byte{}
	spaceFound := false
	for i < len(s) {
		v, size := utf8.DecodeRune(s[i:])
		if unicode.IsSpace(v) {
			spaceFound = true
			i += size
			continue
		}
		if spaceFound {
			ans = append(ans, ' ')
			spaceFound = false
			continue
		} else {
			ans = append(ans, s[i:i+size]...)
		}
		i += size
	}
	return ans
}

func main() {
	message := "hola como estas \U0001F603\t\n\n😃 y\t\tbien       \rlol"
	fmt.Printf("message: %s\n", message)
	encodedUTF8 := []byte(message)
	fmt.Printf("byte slice: %v\n", encodedUTF8)
	fmt.Printf("squashed: \"%s\"\n", string(squashSpaces(encodedUTF8)))

}
