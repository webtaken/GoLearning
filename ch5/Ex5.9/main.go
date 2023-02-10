package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := strings.Join(os.Args[1:], " ")
	ans := expand(str, strings.ToLower)
	fmt.Printf("%s\n", ans)
}

func expand(s string, f func(string) string) string {
	elems := strings.Split(s, " ")
	for i, elem := range elems {
		if strings.HasPrefix(elem, "$") {
			elems[i] = f(elem[1:])
		}
	}
	return strings.Join(elems, " ")
}
