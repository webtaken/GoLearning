package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	dot := strings.LastIndex(s, ".")
	real := ""
	decimal := ""
	if dot >= 0 {
		real = s[:dot]
		decimal = s[dot+1:]
	}
	var buf bytes.Buffer
	iter := 0
	// 12345.45
	for (len(real)-iter)%3 != 0 {
		fmt.Fprintf(&buf, "%s", real[iter:iter+1])
		iter++
	}
	for ; iter < len(real); iter += 3 {
		fmt.Fprintf(&buf, ",%s", real[iter:iter+3])
	}
	ans := ""
	if len(real)%3 == 0 {
		ans = buf.String()[1:]
	} else {
		ans = buf.String()
	}
	if decimal != "" {
		ans += "." + decimal
	}
	return ans
}
func main() {
	ans1 := comma("1254165164354.91951789")
	fmt.Println(ans1)
}
