package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func LockIssue(owner, repo, issueNumber, reason string) error {
	reasonLock := struct {
		LockReason string `json:"lock_reason"`
	}{
		LockReason: reason,
	}

	issueJSON, err := json.Marshal(reasonLock)
	if err != nil {
		return err
	}
	var bearer = "Bearer " + os.Getenv("GITHUB_TOKEN")
	endpoint := fmt.Sprintf("%s/%s/%s/issues/%s/lock", IssuesURL, owner, repo, issueNumber)
	fmt.Println("Endpoint:", endpoint)
	req, _ := http.NewRequest("PUT", endpoint, bytes.NewBuffer(issueJSON))
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/vnd.github+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	// 204 is the status code that github api return when an issue is locked
	if resp.StatusCode != 204 {
		resp.Body.Close()
		return fmt.Errorf("search query failed: %s", resp.Status)
	}
	resp.Body.Close()
	return nil
}
