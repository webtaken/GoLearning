package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"Ex4.13/poster"
)

func main() {

	if len(os.Args[1:]) == 0 {
		log.Fatalf(`Please provide a movie name, e.g.
./bin Guardians of the Galaxy Vol. 2`)
	}
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("%s", err)
	}
	movie, err := poster.RetrieveMovie(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Movie\nTitle: %s\nPoster URL:%s\n", movie.Title, movie.Poster)
	poster.RetrievePoster(movie)
}
