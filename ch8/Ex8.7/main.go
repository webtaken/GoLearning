package main

import (
	"findlinks3/links"
	"flag"
	"fmt"
	"log"
)

type Page struct {
	URL   string
	depth int
}

var websiteURL string
var maxDepth int

func init() {
	flag.IntVar(&maxDepth, "depth", 3, "depth of the crawl (default 3)")
	flag.StringVar(&websiteURL, "url", "https://link-busters.com", "starting url to crawl (default https://link-busters.com)")
}

func main() {
	flag.Parse()

	worklist := make(chan []Page)
	unseenLinks := make(chan Page) // de-duplicated URLS (key) with its depth as value

	initialPage := Page{URL: websiteURL, depth: 1}
	originDomain, err := links.GetDomain(websiteURL)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	var pending int // number of pending sends to worklist
	pending++
	go func() { worklist <- []Page{initialPage} }()

	nCrawlers := 5
	for i := 0; i < nCrawlers; i++ {
		go func() {
			for link := range unseenLinks {
				foundPages := crawl(link, originDomain)
				go func() {
					worklist <- foundPages
					pending--
				}()
			}
		}()
	}
	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for ; pending > 0; pending-- {
		listPages := <-worklist
		for _, page := range listPages {
			if !seen[page.URL] && page.depth <= maxDepth {
				seen[page.URL] = true
				pending++
				unseenLinks <- page
			}
		}
	}
}

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(page Page, domain string) []Page {
	tokens <- struct{}{} // acquire a token
	list, doc, err := links.Extract(page.URL, domain)
	if err != nil {
		log.Print(err)
	}
	<-tokens // release the token
	newPages := make([]Page, len(list))
	for i, url := range list {
		newPages[i] = Page{URL: url, depth: page.depth + 1}
	}
	// Writing mirrored page
	filepath, err := links.GetPathFromURL(page.URL)
	fmt.Printf("Crawled URL: %s\tdepth: %d\n", page.URL, page.depth)
	fmt.Printf("writing page to file: %s\n", filepath)
	if err != nil {
		log.Print(err)
		return newPages
	}
	err = links.WriteHTMLToFile(domain+filepath, doc)
	if err != nil {
		log.Print(err)
	}

	return newPages
}
