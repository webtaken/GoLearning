// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "mapping: %v\n", err)
		os.Exit(1)
	}
	elements := make(map[string]uint32)
	fmt.Printf("Starting the elements mapping of %s...", url)
	mapping(&elements, doc)
	fmt.Printf("Results:\n")
	for tag, n := range elements {
		fmt.Printf("Tag: %9.9s\tCount: %d\n", tag, n)
	}
}

// visit appends to links each link found in n and returns the result.
func mapping(elements *map[string]uint32, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		(*elements)[n.Data]++
	}
	mapping(elements, n.FirstChild)
	mapping(elements, n.NextSibling)
}
