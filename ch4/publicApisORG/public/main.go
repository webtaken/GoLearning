// Package public provides a Go API for the PublicAPIs tracker.
// See https://api.publicapis.org/.
package public

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const PublicAPIsURL = "https://api.publicapis.org/entries"

type APIsResults struct {
	Count   int `json:"count"`
	Entries []*API
}
type API struct {
	Title       string
	Description string
	Auth        string
	HTTPS       bool
	Cors        string
	Link        string
	Category    string
}

// SearchAPIs queries the public API tracker.
func SearchAPIs(terms []string) (*APIsResults, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	fmt.Printf("Query is: %s\n", PublicAPIsURL+"?title="+q)
	resp, err := http.Get(PublicAPIsURL + "?title=" + q)
	if err != nil {
		return nil, err
	}
	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result APIsResults
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
