package movie

import (
	"testing"
	"time"
)

func Test_isShowingAfter(t *testing.T) {
	time1, _ := time.Parse(dateFormat, "11:00:00-00:00")
	time2, _ := time.Parse(dateFormat, "12:00:00-00:00")
	time3, _ := time.Parse(dateFormat, "10:59:59-00:00")
	isShowingAfter11 := isShowingAfter(time1)

	m1 := Movie{Showings: []ShowingTime{ShowingTime{time2}}}
	m2 := Movie{Showings: []ShowingTime{ShowingTime{time3}}}

	if isShowingAfter11(Movie{}) {
		t.Errorf("isShowingAfter11(%s) should be true", time2.String())
	}
	if !isShowingAfter11(m1) {
		t.Errorf("isShowingAfter11(%s) should be true", time2.String())
	}
	if isShowingAfter11(m2) {
		t.Errorf("isShowingAfter11(%s) should be false", time2.String())
	}
}

func Test_isGenre(t *testing.T) {
	isAnimation := isGenre("Animation")
	if isAnimation(Movie{}) {
		t.Error("Should not be an animation")
	}
	if isAnimation(Movie{Genres: []string{"NO", "right"}}) {
		t.Error("Should not be an animation")
	}
	if !isAnimation(Movie{Genres: []string{"YES", "aNimation"}}) {
		t.Error("Should be an animation")
	}
}

func TestMovies_filter(t *testing.T) {
	movies := Movies{Movie{}, Movie{}}
	alwaysInFilter := func(movie Movie) (isIn bool) {
		return true
	}
	allMovies := movies.filter([]movieFilter{alwaysInFilter, alwaysInFilter})

	if len(allMovies) != 2 {
		t.Error("All movies should contain all movies")
	}

	alwaysOutFilter := func(movie Movie) (isIn bool) {
		return false
	}
	allMovies = movies.filter([]movieFilter{alwaysInFilter, alwaysOutFilter})

	if len(allMovies) != 0 {
		t.Error("All movies should have been filtered")
	}
}
