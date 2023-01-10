package index

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const ComicURL = "https://xkcd.com"
const indexFilename = "xkcd_index.txt"

type Comic struct {
	Title      string `json:"title"`
	Link       string `json:"link"`
	Transcript string `json:"transcript"`
	Img        string `json:"img"`
}

func GenerateIndex() {
	file_index, err := os.Create(indexFilename)
	if err != nil {
		log.Fatal(err)
	}

	defer file_index.Close()

	for i := 1; ; i++ {
		resp, err := http.Get(fmt.Sprintf("%s/%s/info.0.json", ComicURL, strconv.Itoa(i)))
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			break // break the loop when we do not receive a correct response
		}
		var resultComic Comic
		if err := json.NewDecoder(resp.Body).Decode(&resultComic); err != nil {
			resp.Body.Close()
			log.Fatal(err)
		}
		resp.Body.Close()

		// Making a cleaning to the data
		resultComic.Transcript = strings.ReplaceAll(resultComic.Transcript, "\n", "")
		resultComic.Transcript = strings.ReplaceAll(resultComic.Transcript, "[", "")
		resultComic.Transcript = strings.ReplaceAll(resultComic.Transcript, "]", "")

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
}
