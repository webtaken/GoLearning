package main

import (
	"fmt"
	"sync"
)

type withdrawOp struct {
	amount int
	okChan chan bool
}

var deposits = make(chan int) // Send amount to deposit
var balances = make(chan int) // Receive balance
var withdraws = make(chan withdrawOp)
var done = make(chan struct{}) // Signal to terminate goroutines

func Deposit(amount int) { deposits <- amount }
func Withdraw(amount int) bool {
	newWithdraw := withdrawOp{
		amount: amount,
		okChan: make(chan bool),
	}
	defer close(newWithdraw.okChan)
	withdraws <- newWithdraw
	return <-newWithdraw.okChan
}
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
		case withdraw := <-withdraws:
			if balance-withdraw.amount < 0 {
				withdraw.okChan <- false
				break
			}
			balance -= withdraw.amount
			withdraw.okChan <- true
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

	wg.Wait() // Wait for all goroutines to finish

	fmt.Printf("Account balance: %d\n", bankAmount)

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			ok := Withdraw(200)
			bankAmount = Balance()
			if !ok {
				fmt.Printf("Couldn't withdraw 200 bills, your current balance is: %d\n",
					bankAmount)
			}
			wg.Done()
		}()
	}

	wg.Wait()   // Wait for all goroutines to finish
	close(done) // Signal the teller to terminate

	fmt.Printf("Account balance: %d\n", bankAmount)
}
