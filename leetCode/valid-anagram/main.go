package main

import "fmt"

func main() {
	s, t := "anagram", "nagaram"
	fmt.Println(isAnagram(s, t))
	s, t = "car", "rac"
	fmt.Println(isAnagram(s, t))
	s, t = "sdaa", "sadd"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[rune]int, 0)
	tMap := make(map[rune]int, 0)
	for i := 0; i < len(s); i++ {
		sMap[rune(s[i])]++
		tMap[rune(t[i])]++
	}
	for key, sVal := range sMap {
		tVal, ok := tMap[key]
		if !ok {
			return false
		} else if tVal != sVal {
			return false
		}
	}
	return true
}
