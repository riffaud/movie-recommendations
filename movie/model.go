package movie

import (
	"encoding/json"
	"strings"
	"time"
)

const (
	dateFormat = "15:04:05-07:00"
)

// Movie data, with a specific showing time format
type Movie struct {
	Name     string
	Rating   int
	Genres   []string
	Showings []ShowingTime
}

// Movies describes a collection of movies
type Movies []Movie

// ShowingTime inherits time
type ShowingTime struct {
	time.Time
}

// UnmarshalJSON defines how showing time gets unmarshalled (time is getting parsed from specific format)
func (j *ShowingTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	st, err := parseTime(s)
	if err != nil {
		return err
	}
	*j = ShowingTime{st}
	return nil
}

// MarshalJSON nothing to do there
func (j ShowingTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}

// DisplayNextShowing is a rendering method to display a Movie next showing
func (m Movie) DisplayNextShowing(t time.Time) string {
	d := m.Name + ", "
	for _, ms := range m.Showings {
		if ms.After(t) {
			return d + "showing at " + ms.Format("3pm")
		}
	}
	return d + " is not showing after " + t.Format("3pm")
}

func parseTime(s string) (time.Time, error) {
	t, err := time.Parse(dateFormat, s)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
