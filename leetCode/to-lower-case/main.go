package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "heLLo"

	fmt.Println(toLowerCase(s))
}

func toLowerCase(s string) string {
	return strings.ToLower(s)
}
