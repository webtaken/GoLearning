package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(a))
	a = []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(a))
}

func productExceptSelf(nums []int) []int {
	prodLeft, prodRight := make([]int, len(nums)), make([]int, len(nums))
	prodLeft[0] = nums[0]
	prodRight[len(nums)-1] = nums[len(nums)-1]
	for i, j := 1, len(nums)-2; i < len(nums); i, j = i+1, j-1 {
		prodLeft[i] = prodLeft[i-1] * nums[i]
		prodRight[j] = prodRight[j+1] * nums[j]
	}
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		r, l := 1, 1
		if i-1 >= 0 {
			r = prodLeft[i-1]
		}
		if i+1 < len(nums) {
			l = prodRight[i+1]
		}
		ans[i] = r * l
	}
	return ans
}
