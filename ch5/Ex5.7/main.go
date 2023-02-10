package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

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
		fmt.Fprintf(os.Stderr, "Ex5.7: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		var attribute string
		for _, attr := range n.Attr {
			attribute += fmt.Sprintf("%s=%q ", attr.Key, attr.Val)
		}
		child := ""
		// <div /> is illegal
		if n.Data == "img" && n.FirstChild == nil {
			child = " /"
		}

		if len(attribute) > 1 {
			attribute = attribute[:len(attribute)-1] // to delete the final space
			fmt.Printf("%*s<%s %s>\n", depth*2, "", n.Data, attribute)
		} else {
			fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, child)
		}
		depth++
	}
	if n.Type == html.TextNode || n.Type == html.CommentNode {
		if !(n.Parent.Type == html.ElementNode &&
			(n.Parent.Data == "script" || n.Parent.Data == "style")) {
			if n.Type == html.CommentNode {
				fmt.Printf("%*s<!--\n", depth*2, "")
			}
			for _, line := range strings.Split(n.Data, "\n") {
				line = strings.TrimSpace(line)
				if line != "" && line != "\n" {
					fmt.Printf("%*s%s\n", depth*2, "", line)
				}
			}
		}
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		// <div /> is illegal
		if !(n.Data == "img" && n.FirstChild == nil) {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
	if n.Type == html.CommentNode {
		fmt.Printf("%*s-->\n", depth*2, "")
	}
}
