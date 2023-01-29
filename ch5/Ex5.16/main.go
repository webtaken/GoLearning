package main

import "fmt"

func main() {
	fmt.Printf("%s\n", join(",", "hola", "como", "estas"))
	fmt.Printf("%s\n", join(","))
	fmt.Printf("%s\n", join("-", "4561", "8498", "2132", "5410"))
}

func join(sep string, strs ...string) string {
	final := ""
	for i, str := range strs {
		final += str
		if i != len(strs)-1 {
			final += sep
		}
	}
	return final
}
