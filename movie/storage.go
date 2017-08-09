package movie

import (
	"encoding/json"
	"io"
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

// StorageFromReader builds a static storage from json data fetched from a reader
func StorageFromReader(r io.Reader) (Storage, error) {
	res := Movies{}
	err := json.NewDecoder(r).Decode(&res)
	return StaticStorage{res}, err
}
