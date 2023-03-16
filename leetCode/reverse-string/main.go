package main

import "fmt"

func main() {
	var str = []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(str)
	fmt.Println(string(str))
	str = []byte{'h', 'a', 'n', 'n', 'a', 'h'}
	fmt.Println(string(str))
}

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
