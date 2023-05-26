package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
	nums = []int{1}
	k = 1
	fmt.Println(topKFrequent(nums, k))
	nums = []int{1, 2, 45, 500, -500, 500, -500}
	k = 2
	fmt.Println(topKFrequent(nums, k))
}

type Pair struct {
	key, val int
}

func topKFrequent(nums []int, k int) []int {
	mapNums := make(map[int]int, 0)
	for _, val := range nums {
		mapNums[val]++
	}
	revMap := make([]Pair, 0)
	for k, v := range mapNums {
		revMap = append(revMap, Pair{key: v, val: k})
	}
	sort.SliceStable(revMap, func(i, j int) bool {
		return revMap[i].key > revMap[j].key
	})
	ans := make([]int, 0)
	for i := 0; i < k; i++ {
		ans = append(ans, revMap[i].val)
	}
	return ans
}
