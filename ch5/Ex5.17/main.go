package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]
	tags := os.Args[2:]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ex5.17: %v\n", err)
		os.Exit(1)
	}
	images := ElementsByTagName(doc, tags...)
	fmt.Printf("%v\n", images)
}

func ElementsByTagName(doc *html.Node, names ...string) []*html.Node {
	var results []*html.Node
	isInNames := func(tag string) bool {
		for _, name := range names {
			if tag == name {
				return true
			}
		}
		return false
	}

	search(&results, doc, isInNames)
	return results
}

func search(ans *[]*html.Node, n *html.Node, f func(tag string) bool) {
	if n.Type == html.ElementNode && f(n.Data) {
		*ans = append(*ans, n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		search(ans, c, f)
	}
}
