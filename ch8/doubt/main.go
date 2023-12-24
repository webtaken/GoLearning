package main

import (
	"fmt"
	"time"
)

func monitor(x1 chan int, x2, c chan struct{}) {
outer:
	for {
		select {
		case x1 <- 0:
			fmt.Printf("x1 is waiting for a message\n")
		case x := <-x1:
			fmt.Printf("x1 received a message %d\n", x)
		case <-x2:
			fmt.Printf("x2 received a message\n")
		case <-c:
			fmt.Printf("c received a message\n")
			fmt.Printf("Closing all...\n")
			break outer
		}
	}

	close(x1)
	close(x2)
	close(c)
}

func main() {
	x, y, c := make(chan int), make(chan struct{}), make(chan struct{})
	go monitor(x, y, c)

	go func() {
		time.Sleep(2 * time.Second)
		x <- 1
	}()
	<-x

	time.Sleep(3 * time.Second)
	y <- struct{}{}

	time.Sleep(5 * time.Second)
	c <- struct{}{}
	t := <-c
	fmt.Println(t)
}
