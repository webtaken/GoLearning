package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"public_apis/public"
)

func main() {
	result, err := public.SearchAPIs(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	jsonResult, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatalf("Error during json decoding %s\n", err)
	}
	fmt.Printf("Found %d APIs:\n", result.Count)
	fmt.Printf("%s\n", jsonResult)
}
