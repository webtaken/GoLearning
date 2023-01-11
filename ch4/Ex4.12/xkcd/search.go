package xkcd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func SearchComic(query string) {
	fmt.Printf("Searching the query...\n")
	idxFile, err := os.Open(IndexFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer idxFile.Close()

	scanner := bufio.NewScanner(idxFile)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	i := 1
	for scanner.Scan() {
		i++
		var comic Comic
		comicStr := scanner.Text()
		if err := json.Unmarshal([]byte(comicStr), &comic); err != nil {
			fmt.Printf("JSON unmarshaling failed on comic %d: %s", i, err)
			continue
		}
		if strings.Contains(comic.Transcript, query) {
			fmt.Printf("################################\n")
			fmt.Printf("Comic %d match your result 😃\n\n", i)
			link := fmt.Sprintf("%s/%d", ComicURL, i)
			fmt.Printf("Check the following link: %s\n", link)
			fmt.Printf("Here is the transcript: %s\n", comic.Transcript)
			fmt.Printf("################################\n\n")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
