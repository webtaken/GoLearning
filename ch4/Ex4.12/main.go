package main

import (
	"fmt"
	"log"
	"os"

	"Ex4.12/xkcd"
)

const minimunRequiredParameters = 2
const usageString = `Usage:
- ./bin <OPTION> <QUERY>
Parameters:
- OPTION: <S|G>, S for searching on the index, G for generating a new comic index.
- QUERY: any string; for G option is optional to add "F" letter to force the generation of the index; for S option is the search term.
	`

func checkIndexExists() bool {
	if _, err := os.Stat(xkcd.IndexFilename); err == nil {
		return true
	}
	return false
}

func main() {
	if len(os.Args) < minimunRequiredParameters {
		log.Fatal(usageString)
	}
	option := os.Args[1]
	if option != "S" && option != "G" {
		log.Fatal(usageString)
	}

	if option == "G" {
		if checkIndexExists() {
			if len(os.Args[2:]) > 0 {
				if os.Args[2] == "F" {
					xkcd.GenerateIndex()
				} else {
					log.Fatal(usageString)
				}
			}
			fmt.Printf("Index already exist please type \"S\" option or force the generation of a new index\n\n%s\n", usageString)
		} else {
			xkcd.GenerateIndex()
		}
	} else {
		if !checkIndexExists() {
			log.Fatalf("Please first generate the index of comics.\n\n%s\n", usageString)
		}
		if len(os.Args[2:]) == 0 {
			log.Fatalf("Please provide a query\n\n%s\n", usageString)
		}
		xkcd.SearchComic(os.Args[2])
	}

}
