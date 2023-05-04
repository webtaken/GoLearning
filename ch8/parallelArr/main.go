package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}
	var wg sync.WaitGroup
	fmt.Printf("Before adding: %v\n", arr)
	for i := range arr {
		wg.Add(1)
		go addOne(arr, i, &wg)
	}
	wg.Wait()
	fmt.Printf("After adding: %v\n", arr)
}

func addOne(arr []int, i int, w *sync.WaitGroup) {
	arr[i]++
	w.Done()
}
