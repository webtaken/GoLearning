// Package links provides a link-extraction function.
package links

import (
	"fmt"
	"ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func URLToSlug(url string) string {
	// Remove the scheme from the URL, if it exists.
	if strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://") {
		url = url[len("https://"):]
	}

	// Remove the query parameters from the URL, if it exists.
	if strings.Contains(url, "?") {
		url = url[:strings.Index(url, "?")]
	}

	// Convert the URL to lowercase.
	url = strings.ToLower(url)

	// Replace all non-alphanumeric characters with hyphens (`-`).
	url = strings.ReplaceAll(url, `[^a-zA-Z0-9]+`, "-")

	// Remove any leading or trailing hyphens.
	url = strings.Trim(url, "-")

	return url
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

func WriteHTML(filename string, node *html.Node) error {
	fileWriter, err := ioutil.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	fmt.Fprintf(fileWriter, "%v", node)
	defer fileWriter.Close()
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
					continue // ignore bad URLs
				}
				if linkDomain != domain {
					// not domain URL now has the URL of the mirorred URL
					a.Val = resp.Request.URL.String()
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
