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
	storage, err = movie.StorageFromURL("http://pastebin.com/raw/cVyp3McN")

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
		s, err := time.ParseInLocation("15:04", *showing, time.Local) // parse time in current timezone
		if err != nil {
			fmt.Println("Error Parsing date, should use format: 15:04")
			return
		}
		s = s.Add(-time.Minute * 30)
		searchParams.Showing = s
	}

	for _, m := range storage.Load(searchParams) {
		fmt.Println(m.DisplayNextShowing(searchParams.Showing))
	}
}
