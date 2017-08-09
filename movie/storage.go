package movie

import (
	"encoding/json"
	"net/http"
	"time"
)

// Storage type to describe a storage that can load movies
type Storage interface {
	Load(params SearchParams) Movies
}

// StaticStorage is a sort of storage that contains all movies in memory
type StaticStorage struct {
	Movies Movies
}

// SearchParams to query/search a storage
type SearchParams struct {
	Genre   string
	Showing time.Time
}

// Load and filter data from a static storage
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

// StorageFromURL builds a static storage from json data fetched by http get
func StorageFromURL(url string) (Storage, error) {
	response, err := http.Get(url)
	if err != nil {
		return StaticStorage{}, err
	}
	defer response.Body.Close()
	res := Movies{}
	json.NewDecoder(response.Body).Decode(&res)
	return StaticStorage{res}, nil
}
