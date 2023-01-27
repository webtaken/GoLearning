package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"Ex5.13/links"
)

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:], os.Args[1])
}

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string, mainDomain string) []string, worklist []string, mainDomain string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item, mainDomain)...)
			}
		}
	}
}

func crawl(url string, mainDomain string) []string {
	if strings.HasPrefix(url, mainDomain) {
		filename := url
		filename = strings.ReplaceAll(filename, "https://", "")
		filename = strings.ReplaceAll(filename, ".html", "")
		filename = strings.ReplaceAll(filename, ".", "_")
		filename = fmt.Sprintf("crawling/page_%s.html", filename)
		f, err := os.Create(filename)
		if err != nil {
			log.Print(err)
		}
		defer f.Close()
		links.ExtractContent(url, f)
	}
	validator := func(link string) bool {
		return strings.HasPrefix(link, mainDomain)
	}
	fmt.Printf("Crawling urls from: %s\n", url)
	list, err := links.Extract(url, validator)
	if err != nil {
		log.Print(err)
	}
	return list
}
