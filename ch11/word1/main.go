// Package word provides utilities for word games.
package main

import (
	"fmt"
	"word1/word"
)

func main() {
	isPalindrome := word.IsPalindrome("abacaba")
	fmt.Println(isPalindrome)
}
