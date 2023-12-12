package main

import (
	"fmt"
	"math/rand"
)

const SIZE = 10000

func main() {
	arr := [SIZE]int{0}

	for i := 0; i < SIZE; i++ {
		// arr[i] = i
		arr[i] = rand.Intn(10)
	}

	fmt.Printf("arr: %v\n", arr[:5])
	total := make(chan int)
	done := make(chan struct{})
	length := SIZE / 100
	for i := 0; i < SIZE; i += length {
		go func(start, end int) {
			sumArr(&arr, start, end, total)
		}(i, i+length)
	}
	go func() {
		defer func() {
			done <- struct{}{}
		}()
		totalPieces := SIZE / length
		piecesCount := 0
		totalSum := 0
		for x := range total {
			totalSum += x
			piecesCount++
			if piecesCount >= totalPieces {
				fmt.Printf("Total Sum is: %d\n", totalSum)
				return
			}
		}
	}()

	<-done
}

func sumArr(arr *[SIZE]int, start, end int, res chan<- int) {
	tmpSum := 0
	for i := start; i < end; i++ {
		tmpSum += (*arr)[i]
	}
	res <- tmpSum
}
