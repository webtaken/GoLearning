package main

import "fmt"

func removeDuplicates(s []string) []string {
	if len(s) <= 1 {
		return s
	}
	i := 0
	ans := []string{}
	for i < len(s)-1 {
		j := i + 1
		for j < len(s) && s[i] == s[j] {
			j++
		}
		ans = append(ans, s[i])
		i = j
	}
	return ans
}

func main() {
	strs := []string{"a", "b", "c", "c", "c", "c", "d", "e", "e"}
	fmt.Printf("before: %v\n", strs)
	fmt.Printf("after: %v\n", removeDuplicates(strs))
}
