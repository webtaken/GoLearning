package main

import "fmt"

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))
	nums = []int{100, 4, 200, 1, 3, 2, 2, 1}
	fmt.Println(longestConsecutive(nums))
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	seq := make(map[int]bool)
	max := 1
	for _, val := range nums {
		seq[val] = true
	}
	for _, val := range nums {
		l := val - 1
		_, ok := seq[l]
		if !ok {
			c := 1
			r := val + 1
			for {
				if _, ok := seq[r]; !ok {
					break
				}
				c++
				r++
			}
			if c > max {
				max = c
			}
		}
	}
	return max
}
