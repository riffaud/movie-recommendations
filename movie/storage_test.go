package movie

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestStorageFromReader(t *testing.T) {
	storage, err := StorageFromReader(strings.NewReader("[]"))

	_, ok := storage.(StaticStorage)
	if err != nil || !ok {
		t.Error("Should have returned a static storage")
	}
}

// Integration test without stubbing the models or the filters
func TestStaticStorage_Load(t *testing.T) {
	time1, _ := time.Parse(dateFormat, "11:00:00-00:00")
	storage, _ := StorageFromReader(strings.NewReader(
		`[
			{
				"name": "Moonlight",
				"rating": 98,
				"genres": [
					"Drama"
				],
				"showings": [
					"18:30:00+00:00",
					"20:30:00+00:00"
				]
			},
			{
				"name": "Movie1",
				"rating": 98,
				"genres": [
					"Comedie"
				],
				"showings": [
					"18:30:00+00:00"
				]
			}
		]`,
	))

	returned := storage.Load(SearchParams{Genre: "Drama", Showing: time1})

	if len(returned) != 1 {
		fmt.Println(returned)
		t.Error("Should have returned 1 result")
	}
}
