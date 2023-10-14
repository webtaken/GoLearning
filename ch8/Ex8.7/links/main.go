// Package links provides a link-extraction function.
package links

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func GetPathFromURL(link string) (string, error) {
	// Parse the URL.
	parsedURL, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	path := parsedURL.Path
	if path == "" {
		return "/index.html", nil
	}
	if path[len(path)-1] == '/' {
		return path + "index.html", nil
	}
	// remove .html suffix
	path = strings.TrimSuffix(path, ".html")
	path += ".html"
	return path, nil
}

func GetDomain(link string) (string, error) {
	// Parse the URL.
	parsedURL, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	// Get the hostname of the URL.
	// Remove the `www.` prefix from the hostname, if it exists.
	hostname := strings.TrimPrefix(parsedURL.Hostname(), "www.")
	return hostname, nil
}

func WriteHTMLToFile(filename string, doc *html.Node) error {
	dir := filepath.Dir(filename)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = html.Render(file, doc)
	if err != nil {
		return err
	}

	return nil
}

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string, domain string) ([]string, *html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				linkDomain, err := GetDomain(link.String())
				if err != nil {
					continue
				}
				if linkDomain != domain {
					// domain not belonging to the original domain
					// fmt.Printf("changing link from %s to %s\n",
					// 	link.String(), resp.Request.URL.String())
					a.Val = resp.Request.URL.String()
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)

	return links, doc, nil
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
