package main

import "fmt"

func main() {
	temp := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temp))
	temp = []int{73}
	fmt.Println(dailyTemperatures(temp))
	temp = []int{73, 72, 71, 72}
	fmt.Println(dailyTemperatures(temp))
}

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	s := make([][2]int, 0)
	for i := 0; i < len(temperatures); i++ {
		for len(s) > 0 && s[len(s)-1][0] < temperatures[i] {
			ans[s[len(s)-1][1]] = i - s[len(s)-1][1]
			s = s[:len(s)-1]
		}
		s = append(s, [2]int{temperatures[i], i})
	}
	return ans
}
