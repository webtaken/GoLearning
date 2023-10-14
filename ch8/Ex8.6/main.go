package main

import (
	"findlinks3/links"
	"flag"
	"fmt"
	"log"
)

var websiteURL string

func init() {
	flag.StringVar(&websiteURL, "url", "https://link-busters.com", "starting url to crawl (default https://link-busters.com)")
}

func main() {
	flag.Parse()

	worklist := make(chan []string)
	unseenLinks := make(chan string) // de-duplicated URLS (key) with its depth as value

	originDomain, err := links.GetDomain(websiteURL)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	go func() { worklist <- []string{websiteURL} }()

	nCrawlers := 5
	for i := 0; i < nCrawlers; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link, originDomain)
				go func() { worklist <- foundLinks }()
			}
		}()
	}
	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

func crawl(url, domain string) []string {
	fmt.Printf("Crawling urls from: %s\n", url)
	list, _, err := links.Extract(url, domain)
	if err != nil {
		log.Print(err)
	}
	// Writing mirrored page
	filename := links.URLToSlug(url)
	fmt.Printf("writing page to file: %s.html\n", filename)
	// err = links.WriteHTML(filename, page)
	// if err != nil {
	// 	log.Print(err)
	// }
	return list
}
