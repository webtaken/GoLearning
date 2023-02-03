package main

import "fmt"

func longestPalindrome(s string) string {
	res := ""
	resLen := 0
	for i := 0; i < len(s); i++ {
		// odd palindromes
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if (r - l + 1) > resLen {
				res = s[l : r+1]
				resLen = r - l + 1
			}
			r++
			l--
		}
		// even palindromes
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if (r - l + 1) > resLen {
				res = s[l : r+1]
				resLen = r - l + 1
			}
			r++
			l--
		}
	}
	return res
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("abbbddeddc"))
}
