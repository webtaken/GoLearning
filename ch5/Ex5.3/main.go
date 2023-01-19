package main

import (
	"bytes"
	"fmt"
	"io"
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
		fmt.Fprintf(os.Stderr, "visit1: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
}

// visit appends to links each link found in n and returns the result.
func visit(n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode && n.Data != "style" && n.Data != "script" {
		result_str := renderNode(n)
		if result_str != "" {
			fmt.Printf("%s", result_str)
		}
	}
	visit(n.FirstChild)
	visit(n.NextSibling)
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	str_tag := buf.String()
	str_tag = strings.ReplaceAll(str_tag, " ", "")
	return str_tag
}
