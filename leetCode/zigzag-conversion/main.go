package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var table = make(map[int]string)
	j, add := 0, 1
	for i := 0; i < len(s); i++ {
		table[j] += string(s[i])
		j += add
		if j == 0 {
			add = 1
		} else if j == (numRows - 1) {
			add = -1
		}
	}
	ans := ""
	for i := 0; i < numRows; i++ {
		ans += table[i]
	}
	return ans
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("PAYPALISHIRING", 2))
	fmt.Println(convert("PAYPALISHIRING", 20))
	fmt.Println(convert("PAYPALISHIRING", 1))
}
