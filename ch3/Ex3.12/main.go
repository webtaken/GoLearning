package main

import (
	"bytes"
	"fmt"
	"os"
)

func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	runeS2 := []byte(s2)

	for i := 0; i < len(s1); i++ {
		j := bytes.LastIndex(runeS2[i:], []byte(s1[i:i+1]))
		if j < 0 {
			return false
		}
		runeS2[i], runeS2[i+j] = runeS2[i+j], runeS2[i]
	}
	return true
}

func main() {
	s1, s2 := os.Args[1], os.Args[2]
	fmt.Printf("A(%s,%s)=%v\n", s1, s2, anagram(s1, s2))
}
