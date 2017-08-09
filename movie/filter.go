package movie

import (
	"strings"
	"time"
)

// Custom function type to filter a movie
type movieFilter func(movie Movie) (isIn bool)

// Returns true when the movie is showing after a given time
func isShowingAfter(st time.Time) movieFilter {
	return func(m Movie) (isIn bool) {
		for _, ms := range m.Showings {
			if ms.After(st) {
				return true
			}
		}
		return false
	}
}

// Returns true when movie has the defined genre
func isGenre(sg string) movieFilter {
	sg = strings.ToLower(sg)
	return func(movie Movie) (isIn bool) {
		for _, g := range movie.Genres {
			if strings.ToLower(g) == sg {
				return true
			}
		}
		return false
	}
}

// Filter a movies based on a collection of filters
func (movies Movies) filter(filters []movieFilter) (ret Movies) {
	for _, m := range movies {
		isIn := true
		for i := 0; isIn && i < len(filters); i++ {
			isIn = isIn && filters[i](m)
		}
		if isIn {
			ret = append(ret, m)
		}
	}
	return ret
}
