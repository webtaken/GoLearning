package poster

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func RetrieveMovie(words []string) (*Movie, error) {
	movieName := url.QueryEscape(strings.Join(words, "+"))
	query := fmt.Sprintf("%s/?t=%s&apikey=%s", OmdbAPIEndpoint, movieName, os.Getenv("omdb_API_Key"))
	resp, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result Movie
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func RetrievePoster(movie *Movie) {
	movieURL := movie.Poster

	if movieURL == "N/A" || movieURL == "" {
		log.Fatalf("Movie hasn't poster url 🙁\n")
	}

	resp, err := http.Get(movieURL)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}

	posterFileName := strings.ReplaceAll(movie.Title, ":", "") + ".jpg"
	posterFile, err := os.Create(posterFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer posterFile.Close()

	_, err = io.Copy(posterFile, resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Poster downloaded!")
	fmt.Printf("Filename created: %q\n", posterFileName)
}
