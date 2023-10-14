package main

import "fmt"

func main() {
	s := "ab#c"
	t := "ad#c"
	fmt.Printf("%t\n", backspaceCompare(s, t))
	s = "ab##"
	t = "c#d#"
	fmt.Printf("%t\n", backspaceCompare(s, t))
	s = "a#c"
	t = "b"
	fmt.Printf("%t\n", backspaceCompare(s, t))
	s = "##ab"
	t = "#a"
	fmt.Printf("%t\n", backspaceCompare(s, t))
}

func backspaceCompare(s string, t string) bool {
	readString := func(n string) string {
		arr := make([]byte, 0)
		for _, n_val := range n {
			if n_val == '#' {
				if len(arr) > 0 {
					arr = arr[:len(arr)-1]
				}
				continue
			}
			arr = append(arr, byte(n_val))
		}
		return string(arr)
	}
	return readString(s) == readString(t)
}
