package main

import (
	"findlinks3/links"
	"fmt"
	"log"
	"os"
)

func main() {
	worklist := make(chan []string)
	// Start with the command-line arguments.
	go func() { worklist <- os.Args[1:] }()
	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Printf("Crawling urls from: %s\n", url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
