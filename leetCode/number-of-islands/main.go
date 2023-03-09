package main

import "fmt"

func main() {
	var grid = [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'1', '0', '1', '0', '1'},
	}

	var grid2 = [][]byte{
		{'0', '1', '1', '0', '1'},
	}

	fmt.Printf("%v\n", numIslands(grid))
	fmt.Printf("%v\n", numIslands(grid2))
}

func numIslands(grid [][]byte) int {
	count := 0
	m, n := len(grid), len(grid[0])
	v := make([][]bool, m)
	for i := range v {
		v[i] = make([]bool, n)
	}

	isInBorder := func(i, j int) bool {
		return 0 <= i && i < m && 0 <= j && j < n
	}

	visited := func(i, j int) bool {
		if isInBorder(i, j) {
			return v[i][j]
		}
		return false
	}

	is_island := func(i, j int) bool {
		if isInBorder(i, j) {
			return grid[i][j] == '1'
		}
		return false
	}

	var bfs func(i, j int) bool
	bfs = func(i, j int) bool {
		is_terra := false
		if !visited(i, j) && isInBorder(i, j) {
			v[i][j] = true
			if is_island(i, j) {
				is_terra = true
				bfs(i-1, j)
				bfs(i, j+1)
				bfs(i+1, j)
				bfs(i, j-1)
			}
		}
		return is_terra
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if bfs(i, j) {
				count++
			}
		}
	}

	return count
}
