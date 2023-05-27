package main

import "fmt"

func main() {
	s := "()"
	fmt.Println(isValid(s))
	s = "()[]{}{}([])"
	fmt.Println(isValid(s))
	s = "(([{(})]))"
	fmt.Println(isValid(s))
	s = "(){}}{"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		if i > 0 && len(stack) > 1 {
			switch stack[len(stack)-1] {
			case ')':
				if stack[len(stack)-2] == '(' {
					stack = stack[:len(stack)-2]
				}
			case '}':
				if stack[len(stack)-2] == '{' {
					stack = stack[:len(stack)-2]
				}
			case ']':
				if stack[len(stack)-2] == '[' {
					stack = stack[:len(stack)-2]
				}
			}
		}
	}
	return len(stack) == 0
}
