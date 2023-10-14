package main

import (
	"fmt"
	"regexp"
	"strings"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}
func reverseString(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	sCopy := strings.ToLower(s)
	sCopy = clearString(sCopy)
	sRev := reverseString(sCopy)
	return sRev == sCopy
}
