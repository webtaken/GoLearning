package main

import (
	"fmt"
	"log"
	"os"
	"packages/popcount"
	"strconv"
)

func main() {
	a, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		log.Fatalf("Error in ParseUint (strconv): %s\n", err)
	}
	fmt.Printf("%v has %v ones\n", a, popcount.PopCountv1(a))
	fmt.Printf("%v has %v ones\n", a, popcount.PopCountv2(a))
	fmt.Printf("%v has %v ones\n", a, popcount.PopCountv3(a))
	fmt.Printf("%v has %v ones\n", a, popcount.PopCountv4(a))
}
