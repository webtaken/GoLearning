package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))
}

func threeSum(nums []int) [][]int {
	var triplets [][]int
	seenTriplet := make(map[[3]int]bool)
	twoSum := func(nums []int, target, secureIdx int) ([][]int, bool) {
		myMap := make(map[int]int)
		var ans [][]int
		okans := false
		for i, num := range nums {
			if i == secureIdx {
				continue
			}
			if _, ok := myMap[target-num]; ok {
				ans = append(ans, []int{myMap[target-num], i})
				okans = true
			}
			myMap[num] = i
		}
		return ans, okans
	}
	for i, num := range nums {
		if tmp, ok := twoSum(nums, -num, i); ok {
			for _, tripletIdx := range tmp {
				triplet := [3]int{num, nums[tripletIdx[0]], nums[tripletIdx[1]]}
				sort.Ints(triplet[:])
				if _, ok2 := seenTriplet[triplet]; !ok2 {
					seenTriplet[triplet] = true
					triplets = append(triplets, triplet[:])
				}

			}
		}
	}
	return triplets
}
