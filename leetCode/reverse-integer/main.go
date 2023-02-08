package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	param, _ := strconv.ParseInt(os.Args[1], 10, 0)
	fmt.Printf("%v\n", reverse(int(param)))
}

func reverse(x int) int {
	str := strconv.FormatInt(int64(x), 10)
	strRev := ""
	for i := len(str) - 1; i >= 0; i-- {
		if string(str[i]) == "-" {
			strRev = "-" + strRev
			continue
		}
		strRev += string(str[i])
	}
	ans, err := strconv.ParseInt(strRev, 10, 32)
	if err != nil {
		return 0
	}
	return int(ans)
}
