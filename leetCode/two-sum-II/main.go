package main

import "fmt"

func main() {
	// arr := []int{2, 7, 11, 5}
	fmt.Println(twoSum([]int{2, 7, 11, 5}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int)
	var ans []int
	for i, num := range nums {
		if _, ok := myMap[target-num]; ok {
			ans = append(ans, myMap[target-num], i+1)
			return ans
		}
		myMap[num] = i + 1
	}
	return ans
}
