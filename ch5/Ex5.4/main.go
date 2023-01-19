// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

type CustomTag struct {
	tag  string
	link string
}

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
		fmt.Fprintf(os.Stderr, "visit: %v\n", err)
		os.Exit(1)
	}
	for _, myTag := range visit(nil, doc) {
		fmt.Println(myTag.tag, ":", myTag.link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []CustomTag, n *html.Node) []CustomTag {
	if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "img" || n.Data == "script") {
		for _, a := range n.Attr {
			if a.Key == "href" || a.Key == "src" {
				links = append(links, CustomTag{tag: n.Data, link: a.Val})
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
