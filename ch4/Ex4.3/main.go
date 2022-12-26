package main

import "fmt"

func reverse(ptr *[10]int) {
	for i, j := 0, 9; i < j; i, j = i+1, j-1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
}

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("%v\n", arr)
	reverse(&arr)
	fmt.Printf("%v\n", arr)
}
