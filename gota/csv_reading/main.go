// Go program to illustrate
// How to read a csv file
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.File
	file, err := os.Open("grouped_titles.csv")

	// Checks for the error
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	// Closes the file
	defer file.Close()

	// The csv.NewReader() function is called in
	// which the object os.File passed as its parameter
	// and this creates a new csv.Reader that reads
	// from the file
	reader := csv.NewReader(file)

	// ReadAll reads all the records from the CSV file
	// and Returns them as slice of slices of string
	// and an error if any
	records, err := reader.ReadAll()

	// Checks for the error
	if err != nil {
		fmt.Println("Error reading records")
	}

	// Loop to iterate through
	// and print each of the string slice
	for i, eachrecord := range records {
		for _, record := range eachrecord {
			fmt.Println(record)
		}
		fmt.Printf("\n")
		if i > 10 {
			break
		}
	}
}
