package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"text/tabwriter"
)

type Mat struct {
	m, n int
	arr  []int
}

func Constructor(m, n int) Mat {
	newMat := Mat{
		m:   m,
		n:   n,
		arr: make([]int, m*n),
	}
	return newMat
}

// Insert a number into the matrix given its position
func (A Mat) Insert(i, j, v int) {
	if i < 0 || j < 0 {
		log.Printf("Indexes out of bounds")
		return
	}
	if i >= A.m || j >= A.n {
		log.Printf("Indexes out of bounds")
		return
	}
	A.arr[i*A.n+j] = v
}

func (A Mat) PrintMat() {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	for i := 0; i < A.m*A.n; i += A.n {
		arr := A.arr[i : i+A.n]
		strRow := ""
		for _, v := range arr {
			strRow += fmt.Sprintf("%d\t", v)
		}
		fmt.Fprintln(w, strRow)
	}
	w.Flush()
}

// V1
// func (A Mat) Add(B Mat) {
// 	if A.m != B.m || A.n != B.n {
// 		log.Printf("Cannot add matrix with different dimensions\n")
// 		return
// 	}
// 	var wg sync.WaitGroup
// 	wg.Add(len(A.arr))
// 	for index := range A.arr {
// 		go func(i int) {
// 			A.arr[i] += B.arr[i]
// 			wg.Done()
// 		}(index)
// 	}
// 	wg.Wait()
// }

// V2
func (A Mat) Add(B Mat) {
	if A.m != B.m || A.n != B.n {
		log.Printf("Cannot add matrix with different dimensions\n")
		return
	}
	var wg sync.WaitGroup
	wg.Add(A.m)
	for index := 0; index < A.m*A.n; index += A.n {
		go func(i, j int) {
			for k := i; k < j; k++ {
				A.arr[k] += B.arr[k]
			}
			wg.Done()
		}(index, index+A.n)
	}
	wg.Wait()
}

func (A Mat) Sub(B Mat) {
	if A.m != B.m || A.n != B.n {
		log.Printf("Cannot subtract matrix with different dimensions\n")
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(A.arr))
	for index := range A.arr {
		go func(i int) {
			A.arr[i] -= B.arr[i]
			wg.Done()
		}(index)
	}
	wg.Wait()
}

func (A Mat) Randomize(n int) {
	for i := 0; i < A.m; i++ {
		for j := 0; j < A.n; j++ {
			A.Insert(i, j, rand.Intn(n))
		}
	}
}

func main() {
	m, n := 5, 5
	A := Constructor(m, n)
	A.Randomize(10)
	A.PrintMat()
	fmt.Println()

	B := Constructor(m, n)
	B.Randomize(10)
	B.PrintMat()
	fmt.Println()

	A.Sub(B)
	A.PrintMat()
	fmt.Println()

	A.Add(B)
	A.PrintMat()
}
