package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	n := 3
	fmt.Println(generateParenthesis(n))
	n = 1
	fmt.Println(generateParenthesis(n))
	n = 8
	fmt.Println(generateParenthesis(n))
}

func generateParenthesis(n int) []string {
	result := []string{}
	binary := func(i int) string {
		ans := ""
		for i != 0 {
			tmp := strconv.Itoa(i % 2)
			ans = ans + tmp
			i /= 2
		}
		return ans
	}
	checker := func(bin string) (string, bool) {
		zeroes := 0
		ones := 0
		for _, val := range bin {
			if val == '0' {
				zeroes++
			} else {
				ones++
			}
		}
		if ones != n || zeroes != n {
			return "", false
		}
		s := []rune{}
		for _, val := range bin {
			s = append(s, val)
			if len(s) >= 2 {
				if s[len(s)-1] == '1' && s[len(s)-2] == '0' {
					s = s[:len(s)-2]
				}
			}
		}
		if len(s) != 0 {
			return "", false
		}
		ans := strings.ReplaceAll(bin, "0", "(")
		ans = strings.ReplaceAll(ans, "1", ")")
		return ans, true
	}

	for i := 0; i < int(math.Pow(2, float64(2*n))); i++ {
		str, ok := checker(binary(i))
		if ok {
			result = append(result, str)
		}
	}
	return result
}
