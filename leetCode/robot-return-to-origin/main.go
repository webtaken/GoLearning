package main

import "fmt"

func main() {
	moves := "UDLRR"
	fmt.Println(judgeCircle(moves))
}

func judgeCircle(moves string) bool {
	pos := [2]int{0, 0}
	for _, move := range moves {
		if move == 'U' {
			pos[1]++
		} else if move == 'D' {
			pos[1]--
		} else if move == 'R' {
			pos[0]++
		} else {
			pos[0]--
		}
	}
	return pos[0] == 0 && pos[1] == 0
}
