package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"Ex8.10/links"
)

type Page struct {
	URL   string
	depth int
}

var websiteURL string
var maxDepth int
var workers int

// tokens is a counting semaphore used to
// enforce a limit of "workers" concurrent requests.
var tokens = make(chan struct{})

func init() {
	flag.IntVar(&maxDepth, "depth", 3, "depth of the crawl (default 3)")
	flag.IntVar(&workers, "workers", 10, "number of workers (default 10)")
	flag.StringVar(&websiteURL, "url", "https://link-busters.com", "starting url to crawl (default https://link-busters.com)")
}

var cancel = make(chan struct{})

func main() {
	flag.Parse()

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(cancel)
	}()

	tokens = make(chan struct{}, workers)

	worklist := make(chan []Page)
	unseenLinks := make(chan Page) // de-duplicated URLS (key) with its depth as value

	var pending int // number of pending sends to worklist
	initialPage := Page{URL: websiteURL, depth: 1}

	pending++
	go func() { worklist <- []Page{initialPage} }()

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for ; pending > 0; pending-- {
		pages := <-worklist
		for _, page := range pages {
			if !seen[page.URL] && page.depth <= maxDepth {
				pending++
				go func(page Page) {
					worklist <- crawl(page)
				}(page)
			}
		}
	}
	close(unseenLinks)
	close(worklist)
}

func crawl(page Page) []Page {
	tokens <- struct{}{} // acquire a token
	fmt.Printf("Crawling %s\tdepth %d\n", page.URL, page.depth)
	list, err := links.Extract(page.URL, cancel)
	if err != nil {
		log.Print(err)
	}
	<-tokens // release the token
	newPages := make([]Page, len(list))
	for i, url := range list {
		newPages[i] = Page{URL: url, depth: page.depth + 1}
	}
	return newPages
}
