package main

import (
	"fmt"

	"Ex11.2/intset"
)

func main() {
	newSet := intset.IntSet{}

	newSet.Add(4)
	fmt.Println(newSet)
}
