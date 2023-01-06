package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func UpdateIssue(owner, repo, issueNumber string, data *Issue) (*Issue, error) {
	issueJSON, err := json.Marshal(*data)
	if err != nil {
		return nil, err
	}
	var bearer = "Bearer " + os.Getenv("GITHUB_TOKEN")
	endpoint := fmt.Sprintf("%s/%s/%s/issues/%s", IssuesURL, owner, repo, issueNumber)
	fmt.Println("Endpoint:", endpoint)
	req, _ := http.NewRequest("PATCH", endpoint, bytes.NewBuffer(issueJSON))
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/vnd.github+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	// 200 is the status code that github api return when an issue is updated
	if resp.StatusCode != 200 {
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
