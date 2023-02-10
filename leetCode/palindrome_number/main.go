package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	strNum := strconv.FormatInt(int64(x), 10)
	revStrNum := ""
	for _, digit := range strNum {
		revStrNum = string(digit) + revStrNum
	}

	return strNum == revStrNum
}

func main() {
	fmt.Println(isPalindrome(505))
	fmt.Println(isPalindrome(-101))
	fmt.Println(isPalindrome(0110))
	fmt.Println(isPalindrome(1007010000111111111))
}
