package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var ch = make(chan string)
	var cancel = make(chan struct{})
	if len(os.Args) <= 2 {
		log.Fatalf("Not enough arguments")
	}

	urls := os.Args[1:]
	for _, url := range urls {
		go func(url string) {
			ch <- GetPage(url, cancel)
		}(url)
	}

	fmt.Println(<-ch)
	close(cancel)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// GetPage returns the content of a web page
func GetPage(url string, cancel <-chan struct{}) string {
	start := time.Now()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Sprint(err)
	}
	req.Cancel = cancel

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Sprint(err)
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		return fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	return fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
