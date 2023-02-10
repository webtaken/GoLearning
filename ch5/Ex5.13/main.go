package main

import (
	"fmt"
	"log"
	"net/url"
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
func breadthFirst(f func(item string) []string, worklist []string, mainDomain string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

var domain string

func crawl(a string) []string {
	if domain == "" {
		p, err := url.Parse(a)
		if err != nil {
			log.Fatalf("crawl %s get: %v", err)
		}
		domain = p.Hostname()
		if strings.HasPrefix(domain, "www.") {
			domain = domain[4:]
		}
		fmt.Printf("Domain: %s \n\n", domain)
	}

	list, err := links.Extract(a)
	if err != nil {
		log.Print(err)
	}

	// filter out all links with different domain
	out := list[:0]
	for _, l := range list {
		p, err := url.Parse(l)
		if err != nil {
			// skip invalid url
			continue
		}
		if strings.Contains(p.Hostname(), domain) {
			fmt.Println(l)
			out = append(out, l)
		}
	}
	return out
}
