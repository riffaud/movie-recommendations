package movie

import (
	"encoding/json"
	"net/http"
	"time"
)

type Movie struct {
	Name     string
	Rating   int
	Genres   []string
	Showings []ShowingTime
}

type Movies []Movie

type Storage interface {
	Load(params SearchParams) Movies
}

type StaticStorage struct {
	Movies Movies
}

type SearchParams struct {
	Genre   string
	Showing time.Time
}

func (s StaticStorage) Load(params SearchParams) (ret Movies) {
	filters := []movieFilter{}

	if params.Genre != "" {
		filters = append(filters, isGenre(params.Genre))
	}
	if params.Showing != (time.Time{}) {
		filters = append(filters, isShowingAfter(params.Showing))
	}

	return s.Movies.filter(filters)
}

func StorageFromUrl(url string) (Storage, error) {
	response, err := http.Get(url)
	if err != nil {
		return StaticStorage{}, err
	}
	defer response.Body.Close()
	res := Movies{}
	json.NewDecoder(response.Body).Decode(&res)
	return StaticStorage{res}, nil
}
