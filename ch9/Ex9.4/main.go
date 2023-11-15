package main

import (
	"fmt"
	"time"
)

func pipeline(stages int) (in chan int, out chan int) {
	out = make(chan int)
	first := out
	for i := 0; i < stages; i++ {
		in = out
		out = make(chan int)
		go func(in chan int, out chan int) {
			for v := range in {
				out <- v
			}
			close(out)
		}(in, out)
	}
	return first, out
}

func main() {
	start := time.Now()
	in, out := pipeline(800000)
	in <- 1
	x := <-out
	fmt.Println(time.Since(start))
	fmt.Println(x)
	start = time.Now()
	in <- 2
	x = <-out
	fmt.Println(time.Since(start))
	fmt.Println(x)
	close(in)
	close(out)
}
