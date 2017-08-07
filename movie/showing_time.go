package movie

import (
	"encoding/json"
	"strings"
	"time"
)

// todo:: unse constants
const (
	dateFormat = "15:04:05-07:00"
)

type ShowingTime struct {
	time.Time
}

// follow the example from the doc to implement custom marshall and unmarshal functions https://golang.org/pkg/encoding/json/#pkg-overview
func (j *ShowingTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	st, err := ParseTime(s)
	if err != nil {
		return err
	}
	*j = ShowingTime{st}
	return nil
}

func (j ShowingTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}

// https://stackoverflow.com/questions/14106541/go-parsing-date-time-strings-which-are-not-standard-formats
func ParseTime(s string) (time.Time, error) {
	t, err := time.Parse(dateFormat, s)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
