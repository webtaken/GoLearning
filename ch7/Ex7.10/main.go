package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		j := s.Len() - i - 1
		if !(!s.Less(i, j) && !s.Less(j, i)) {
			return false
		}
	}
	return true
}

type IntsSlice []int

func (s IntsSlice) Len() int           { return len(s) }
func (s IntsSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s IntsSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	var myInts = []int{1, 2, 3, 5, 3, 2, 1}
	fmt.Println(IsPalindrome(IntsSlice(myInts)))
	var myInts2 = []int{1, 2, 3, 3, 2, 10}
	fmt.Println(IsPalindrome(IntsSlice(myInts2)))
}
