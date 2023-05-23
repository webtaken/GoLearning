package main

import "fmt"

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
	strs = []string{""}
	fmt.Println(groupAnagrams(strs))
	strs = []string{"a"}
	fmt.Println(groupAnagrams(strs))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[rune]int, 0)
	tMap := make(map[rune]int, 0)
	for i := 0; i < len(s); i++ {
		sMap[rune(s[i])]++
		tMap[rune(t[i])]++
	}
	for key, sVal := range sMap {
		tVal, ok := tMap[key]
		if !ok {
			return false
		} else if tVal != sVal {
			return false
		}
	}
	return true
}

func mapping(s string) [26]int {
	ans := [26]int{}
	for i := 0; i < len(s); i++ {
		ans[s[i]-'a']++
	}
	return ans
}

// func groupAnagrams(strs []string) [][]string {
// 	selected := make([]bool, len(strs))
// 	ans := make([][]string, 0)
// 	for i := 0; i < len(strs); i++ {
// 		tmpAnagram := make([]string, 0)
// 		if !selected[i] {
// 			tmpAnagram = append(tmpAnagram, strs[i])
// 		} else {
// 			continue
// 		}
// 		for j := i + 1; j < len(strs); j++ {
// 			if isAnagram(strs[i], strs[j]) {
// 				tmpAnagram = append(tmpAnagram, strs[j])
// 				selected[j] = true
// 			}
// 		}
// 		ans = append(ans, tmpAnagram)
// 		selected[i] = true
// 	}
// 	return ans
// }

func groupAnagrams(strs []string) [][]string {
	mapStrs := make(map[[26]int][]string, 0)
	ans := make([][]string, 0)
	for i := 0; i < len(strs); i++ {
		mapKey := mapping(strs[i])
		mapStrs[mapKey] = append(mapStrs[mapKey], strs[i])
	}
	for _, val := range mapStrs {
		ans = append(ans, val)
	}
	return ans
}
