package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/riffaudo/movie-recommendations/movie"
)

var (
	storage movie.Storage
)

func init() {
	var err error
	storage, err = movie.StorageFromUrl("http://pastebin.com/raw/cVyp3McN")

	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	genre := flag.String("genre", "", "Movie genre to return: Action, Animation")
	showing := flag.String("showing", "", "Showing time, format 10:00")
	flag.Parse()

	searchParams := movie.SearchParams{}

	if *genre != "" {
		searchParams.Genre = *genre
	}

	if *showing != "" {
		s, err := time.Parse("15:04", *showing)
		if err != nil {
			fmt.Println(err)
			return
		}
		searchParams.Showing = s
	}

	fmt.Println(storage.Load(searchParams))
}
