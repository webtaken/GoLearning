package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func CreateIssue(owner, repo, title, body string) (*Issue, error) {
	newIssue := Issue{Title: title, Body: body}
	issueJSON, err := json.Marshal(newIssue)
	if err != nil {
		return nil, err
	}
	var bearer = "Bearer " + os.Getenv("GITHUB_TOKEN")
	endpoint := fmt.Sprintf("%s/%s/%s/issues", IssuesURL, owner, repo)
	fmt.Println("Endpoint", endpoint)
	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(issueJSON))
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/vnd.github+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	// 201 is the status code that github api return when an issue is created
	if resp.StatusCode != 201 {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
