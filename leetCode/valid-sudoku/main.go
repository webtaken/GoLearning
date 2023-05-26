package main

import "fmt"

func main() {
	v := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Println(isValidSudoku(v))
	v = [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Println(isValidSudoku(v))

}

func isValidSudoku(board [][]byte) bool {
	checkRepeatedRow := func(r int) bool {
		nums := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for _, num := range board[r] {
			if num != '.' {
				nums[num-'0']++
				if nums[num-'0'] > 1 {
					return true
				}
			}
		}
		return false
	}
	checkRepeatedCol := func(c int) bool {
		nums := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for _, row := range board {
			num := row[c]
			if num != '.' {
				nums[num-'0']++
				if nums[num-'0'] > 1 {
					return true
				}
			}
		}
		return false
	}
	checkRepeatedQuadrant := func(r, c int) bool {
		nums := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				num := board[r+i][c+j]
				if num != '.' {
					nums[num-'0']++
					if nums[num-'0'] > 1 {
						return true
					}
				}
			}
		}
		return false
	}
	for i := 0; i < len(board); i++ {
		if checkRepeatedRow(i) {
			return false
		}
		if checkRepeatedCol(i) {
			return false
		}
	}
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if checkRepeatedQuadrant(i, j) {
				return false
			}
		}
	}
	return true
}
