package main

import (
	"log"
	"os"
)

const requiredParameters = 3
const usageString = `Usage:
- ./bin <OPTION> <QUERY>
Parameters:
- OPTION: <S|G>, S for searching on the index, G for generating a new comic index.
- QUERY: any string; for G option is the name of the index; for S oprtion is the search term.
	`

func main() {
	if len(os.Args) < requiredParameters {
		log.Fatal(usageString)
	}

	option, query := os.Args[1], os.Args[2]
	if option != "S" && option != "G" {
		log.Fatal(usageString)
	}

	if option == "S" {

	}

}
