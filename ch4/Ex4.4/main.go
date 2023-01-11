package main

import (
	"fmt"
	"log"
)

func rotate(s []int, n int) []int {
	if n > len(s) {
		log.Fatalf("Cannot rotate by a value outside thi bounds 0 <= x <= len(s) \n")
	}
	ans := make([]int, len(s))
	copy(ans, s)
	for i := 0; i < n; i++ {
		ans[i] = s[(n+i)%len(s)]
	}
	for i, j := n, 0; i < len(s); i, j = i+1, j+1 {
		ans[i] = s[j]
	}
	return ans
}

func main() {
	mySlice := []int{1, 2, 3, 4, 6, 7}
	fmt.Printf("%v\n", mySlice)
	fmt.Printf("%v\n", rotate(mySlice, 5))
}
