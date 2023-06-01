package main

import (
	"fmt"
	"strconv"
)

func main() {
	tokens := []string{"2", "3", "/"}
	fmt.Println(evalRPN(tokens))
	tokens = []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(tokens))
	tokens = []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(tokens))
	tokens = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens))
}

func evalRPN(tokens []string) int {
	s := make([]string, 0)
	for _, val := range tokens {
		s = append(s, val)
		if len(s) >= 3 {
			switch s[len(s)-1] {
			case "+":
				a, _ := strconv.Atoi(s[len(s)-3])
				b, _ := strconv.Atoi(s[len(s)-2])
				tmp := strconv.Itoa(a + b)
				s = s[:len(s)-3]
				s = append(s, tmp)
			case "-":
				a, _ := strconv.Atoi(s[len(s)-3])
				b, _ := strconv.Atoi(s[len(s)-2])
				tmp := strconv.Itoa(a - b)
				s = s[:len(s)-3]
				s = append(s, tmp)
			case "*":
				a, _ := strconv.Atoi(s[len(s)-3])
				b, _ := strconv.Atoi(s[len(s)-2])
				tmp := strconv.Itoa(a * b)
				s = s[:len(s)-3]
				s = append(s, tmp)
			case "/":
				a, _ := strconv.Atoi(s[len(s)-3])
				b, _ := strconv.Atoi(s[len(s)-2])
				tmp := strconv.Itoa(a / b)
				s = s[:len(s)-3]
				s = append(s, tmp)
			}
		}
	}
	ans, _ := strconv.Atoi(s[0])
	return ans
}
