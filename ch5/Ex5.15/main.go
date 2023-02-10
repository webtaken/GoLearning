package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	fmt.Printf("%d\n", max(4, -8, 3))
	fmt.Printf("%d\n", min(4, -8, 3))
	fmt.Printf("%d\n", max())
	fmt.Printf("%d\n", min())
}

func max(vals ...int) int {
	maxNun := math.MinInt
	for _, val := range vals {
		if val > maxNun {
			maxNun = val
		}
	}
	if len(vals) == 0 {
		log.Fatalf("Error max: no provided numbers\n")
	}
	return maxNun
}

func min(vals ...int) int {
	minNun := math.MaxInt
	for _, val := range vals {
		if val < minNun {
			minNun = val
		}
	}
	if len(vals) == 0 {
		log.Fatalf("Error min: no provided numbers\n")
	}
	return minNun
}
