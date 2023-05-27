package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	strs := []string{"neet", "code"}
	fmt.Println(encode(strs))
	fmt.Println(decode(encode(strs)))
	strs = []string{"", "code"}
	fmt.Println(encode(strs))
	fmt.Println(decode(encode(strs)))
	strs = []string{"saul", "roja$"}
	fmt.Println(encode(strs))
	fmt.Println(decode(encode(strs)))
}

func encode(strs []string) string {
	codecs := make([]string, len(strs))
	for i, str := range strs {
		codecs[i] = fmt.Sprintf("%d$%s", len(str), str)
	}
	return strings.Join(codecs, "")
}

func decode(str string) []string {
	ans := make([]string, 0)
	for i := 0; i < len(str); {
		lengthStr := ""
		for str[i] != '$' {
			lengthStr += string(str[i])
			i++
		}
		i++
		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			panic(err)
		}
		if i < len(str) {
			decoded := str[i : i+length]
			ans = append(ans, decoded)
			i = i + length
		}
	}
	return ans
}
