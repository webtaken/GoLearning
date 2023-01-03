// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	// count of invalid UTF-8 characters
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		word := in.Text()
		counts[word]++
	}
	if in.Err() != nil {
		fmt.Fprintln(os.Stderr, in.Err())
		os.Exit(1)
	}
	for word, count := range counts {
		fmt.Printf("%-30q\t%d\n", word, count)
	}

}
