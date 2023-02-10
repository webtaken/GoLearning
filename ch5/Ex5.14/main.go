package main

import "fmt"

func main() {
	graph := map[int32][]int32{
		0: {1, 2},
		1: {0, 3},
		2: {0, 3, 4},
		3: {1, 2, 5},
		4: {2, 5},
	}
	breadthFirstSearch(graph)
}

func breadthFirstSearch(graph map[int32][]int32) {
	var queue []int32
	visited := make(map[int32]bool)
	for k := range graph {
		queue = append(queue, k)
		break
	}
	for len(queue) > 0 {
		fmt.Printf("%d ", queue[0])
		visited[queue[0]] = true
		for _, val := range graph[queue[0]] {
			if !visited[val] {
				visited[val] = true
				queue = append(queue, val)
			}
		}
		queue = queue[1:]
	}
	fmt.Println()
}
