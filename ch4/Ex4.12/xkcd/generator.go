package xkcd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const IndexFilename = "xkcd_index.txt"

func GenerateIndex() {
	const maxNumberComics = 2722
	fmt.Printf("Generating index...\n")
	file_index, err := os.Create(IndexFilename)
	if err != nil {
		log.Fatal(err)
	}

	defer file_index.Close()

	for i := 1; i <= maxNumberComics; i++ {
		resp, err := http.Get(fmt.Sprintf("%s/%s/info.0.json", ComicURL, strconv.Itoa(i)))
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			fmt.Printf("Comic %d not found...\n", i)
			// continue with the requests
			continue
		}
		fmt.Printf("Writing comic %d to the index...\n", i)
		var resultComic Comic
		if err := json.NewDecoder(resp.Body).Decode(&resultComic); err != nil {
			resp.Body.Close()
			log.Fatal(err)
		}
		resp.Body.Close()

		jsonResultComic, err := json.Marshal(resultComic)
		if err != nil {
			log.Fatal(err)
		}
		// Now we write the content of the file to the index
		_, err = file_index.WriteString(fmt.Sprintf("%s\n", string(jsonResultComic)))
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Finished the index generation check the file %s...\n", IndexFilename)
}
