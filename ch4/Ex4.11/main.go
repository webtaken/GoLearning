// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"Ex4.11/github"
	"github.com/joho/godotenv"
)

func readCredentials(repo, owner *string) {
	fmt.Printf("Write the name of the repo: ")
	fmt.Scan(repo)
	fmt.Printf("Write the name of the repo's owner: ")
	fmt.Scan(owner)
}

func console(crudOption string) {
	in := bufio.NewReader(os.Stdin)
	var repo, owner string
	switch crudOption {
	case "create":
		readCredentials(&repo, &owner)
		fmt.Printf("Write the title of the issue: ")
		title, err := in.ReadString('\n')
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}
		fmt.Printf("What's the issue?: ")
		body, err := in.ReadString('\n')
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}
		result, err := github.CreateIssue(owner, repo, title, body)
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}
		fmt.Printf("Issue created successfully 😄\n")
		jsonResult, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}
		fmt.Printf("%s\n", jsonResult)
	case "read":
		var issueNumber string
		readCredentials(&repo, &owner)
		fmt.Printf("Write the issue number: ")
		fmt.Scan(&issueNumber)
		result, err := github.ReadIssue(owner, repo, issueNumber)
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}
		fmt.Printf("Issue readed successfully 😄\n")
		jsonResult, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}
		fmt.Printf("%s\n", jsonResult)
	case "update":
		var issueNumber string
		readCredentials(&repo, &owner)
		fmt.Printf("Write the issue number: ")
		fmt.Scan(&issueNumber)
		newIssue := github.Issue{}
		fmt.Printf("Write the new title of the issue: ")
		title, err := in.ReadString('\n')
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}
		newIssue.Title = title
		fmt.Printf("Write the new body of the issue: ")
		body, err := in.ReadString('\n')
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}
		newIssue.Body = body
		result, err := github.UpdateIssue(owner, repo, issueNumber, &newIssue)
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}
		fmt.Printf("Issue updated successfully 😄\n")
		jsonResult, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}
		fmt.Printf("%s\n", jsonResult)
	case "lock":
		var issueNumber, lockReason string
		readCredentials(&repo, &owner)
		fmt.Printf("Write the issue number: ")
		fmt.Scan(&issueNumber)
		for {
			fmt.Printf(
				`Select one of the following options for locking this issue:
1) "off-topic"
2) "too heated"
3) "resolved"
4) "spam"
`)
			fmt.Scan(&lockReason)
			if lockReason == "1" {
				lockReason = "off-topic"
				break
			} else if lockReason == "2" {
				lockReason = "too heated"
				break
			} else if lockReason == "3" {
				lockReason = "resolved"
				break
			} else if lockReason == "4" {
				lockReason = "spam"
				break
			} else {
				fmt.Printf("Please Select a valid option!\n")
			}
		}

		err := github.LockIssue(owner, repo, issueNumber, lockReason)
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}
		fmt.Printf("Issue Locked successfully 😄\n")
	default:
		log.Fatalf("Choose a valid option: \"create\"|\"read\"|\"update\"|\"lock\"")
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage:\n -./bin <CRUL>\n--------------------------------\n - <CRUD>: e.g. \"create\"|\"read\"|\"update\"|\"lock\"")
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("%s", err)
	}

	console(os.Args[1])
}
