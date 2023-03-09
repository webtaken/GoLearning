package main

import "fmt"

func main() {
	nums := make([]int, 0)
	nums = append(nums, 1, 2, 3, 1)
	fmt.Println(containsDuplicate(nums))
	nums2 := make([]int, 0)
	nums2 = append(nums2, 1, 2, 3, 4)
	fmt.Println(containsDuplicate(nums2))

}

func containsDuplicate(nums []int) bool {
	duplicates := make(map[int]bool, 0)
	for _, val := range nums {
		if _, ok := duplicates[val]; ok {
			return true
		}
		duplicates[val] = true
	}
	return false
}
