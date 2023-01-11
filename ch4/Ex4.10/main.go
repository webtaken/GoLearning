// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"

	"jsonTutorial/github"
)

func main() {
	if len(os.Args[1:]) < 5 {
		log.Fatalf("Usage:\n -./bin <query> <day> <month> <year> <order>\n--------------------------------\n -query: e.g.\"docker\"\n -day[01-(28...31)]: e.g.\"01\"\n -month[01-12]: e.g.\"11\"\n -year: e.g.\"2022\"\n -order: \"before\" or \"after\"\n\nNOTE: Do not consider quotes \"\" in the examples.")
	}
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %55.55s %10.10s\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}
