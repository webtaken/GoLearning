package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]
	id := os.Args[2]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ex5.8: %v\n", err)
		os.Exit(1)
	}
	el := ElementById(doc, id)
	if el != nil {
		fmt.Printf("Found element:\n")
		printNode(el)
	} else {
		fmt.Printf("Not found element.\n")
	}
}

func printNode(node *html.Node) {
	attributes := ""
	for _, attr := range node.Attr {
		attributes += fmt.Sprintf("%s=%q ", attr.Key, attr.Val)
	}
	fmt.Printf("<%s %s/>\n", node.Data, attributes)
}

func ElementById(doc *html.Node, id string) *html.Node {
	return forEachNode(id, doc, findElement)
}

func forEachNode(id string, n *html.Node, search func(n *html.Node, id string) bool) *html.Node {
	if search != nil {
		if search(n, id) {
			return n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if node := forEachNode(id, c, search); node != nil {
			return node
		}
	}

	return nil
}

func findElement(n *html.Node, id string) bool {
	for _, attr := range n.Attr {
		if attr.Key == "id" && attr.Val == id {
			return true
		}
	}
	return false
}

// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
// package main

// import (
// 	"net/http"
// 	"os"

// 	"fmt"

// 	"golang.org/x/net/html"
// )

// func main() {
// 	for _, url := range os.Args[1:] {
// 		outline(url)
// 	}
// }

// func outline(url string) error {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	doc, err := html.Parse(resp.Body)
// 	if err != nil {
// 		return err
// 	}
// 	//!+call
// 	if node := ElementByID(doc, "mainNav"); node != nil {
// 		fmt.Println(node.Type, node.Data)
// 	}
// 	//!-call

// 	return nil
// }

// func ElementByID(doc *html.Node, id string) *html.Node {
// 	pre := func(doc *html.Node) bool {
// 		for _, attr := range doc.Attr {
// 			if attr.Key == "id" && attr.Val == id {
// 				return true
// 			}
// 		}
// 		return false
// 	}

// 	if node, found := forEachNode(doc, pre, nil); found {
// 		return node
// 	}
// 	return nil
// }

// // !+forEachNode
// // forEachNode calls the functions pre(x) and post(x) for each node
// // x in the tree rooted at n. Both functions are optional.
// // pre is called before the children are visited (preorder) and
// // post is called after (postorder).
// func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) (*html.Node, bool) {
// 	if pre != nil {
// 		if pre(n) {
// 			return n, true
// 		}
// 	}

// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		if node, ok := forEachNode(c, pre, post); ok {
// 			return node, true
// 		}
// 	}

// 	if post != nil {
// 		if post(n) {
// 			return n, true
// 		}
// 	}
// 	return nil, false
// }
