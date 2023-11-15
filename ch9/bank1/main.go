package main

import (
	"fmt"
	"sync"
)

var deposits = make(chan int)  // Send amount to deposit
var balances = make(chan int)  // Receive balance
var done = make(chan struct{}) // Signal to terminate goroutines

func Deposit(amount int) { deposits <- amount }

func Balance() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case <-done:
			return
		}
	}
}

func init() {
	go teller() // Start the monitor goroutine
}

func main() {
	numGoroutines := 5
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	bankAmount := 0

	for i := 0; i < numGoroutines; i++ {
		go func() {
			Deposit(100)
			bankAmount = Balance()
			wg.Done()
		}()
	}

	wg.Wait()   // Wait for all goroutines to finish
	close(done) // Signal the teller to terminate

	fmt.Printf("Account balance: %d\n", bankAmount)
}
