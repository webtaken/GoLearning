package main

import "fmt"

func main() {
	s := "hello world \U0001f602 😄"
	fmt.Printf("%v\n %d\n", s, len(s))

	goUsage :=
		`Go is a tool  for managing Go source code.
Usage:
	go command [arguments]`
	fmt.Printf("%v\n", goUsage)
}
