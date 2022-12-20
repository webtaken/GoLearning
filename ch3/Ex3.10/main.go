package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	iter := 0
	// 12345
	for (len(s)-iter)%3 != 0 {
		fmt.Fprintf(&buf, "%s", s[iter:iter+1])
		iter++
	}
	for ; iter < len(s); iter += 3 {
		fmt.Fprintf(&buf, ",%s", s[iter:iter+3])

	}
	if len(s)%3 == 0 {
		return buf.String()[1:]
	}
	return buf.String()
}

func main() {
	ans := comma("12324567")
	fmt.Printf("%s\n", ans)
}
