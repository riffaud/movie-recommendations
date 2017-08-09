package movie

import (
	"testing"
	"time"
)

func TestShowingTime_UnmarshalJSON_Error(t *testing.T) {
	var showingTime1 ShowingTime
	err := showingTime1.UnmarshalJSON([]byte("\"sdadas\""))
	if err == nil {
		t.Error("Expected error when parsing wrong showing time format")
	}
}

func TestShowingTime_UnmarshalJSON(t *testing.T) {
	time1, _ := time.Parse(dateFormat, "11:00:00-00:00")
	expectedShowingTime := ShowingTime{time1}
	var showingTime1 ShowingTime

	err := showingTime1.UnmarshalJSON([]byte("\"11:00:00-00:00\""))
	if err != nil {
		t.Error("Valid showing time should have been parsed")
	}
	if showingTime1.String() != expectedShowingTime.String() {
		t.Errorf("UnmarshalJSON expected showing time `%s`, found `%s`", expectedShowingTime.String(), showingTime1.String())
	}
}

func TestMovie_DisplayNextShowing(t *testing.T) {
	time1, _ := time.Parse(dateFormat, "11:00:00-00:00")
	time2, _ := time.Parse(dateFormat, "12:00:00-00:00")
	time3, _ := time.Parse(dateFormat, "13:00:00-00:00")

	movie := Movie{
		Name: "Movie",
		Showings: []ShowingTime{
			ShowingTime{time2},
			ShowingTime{time3},
		},
	}

	showing := movie.DisplayNextShowing(time1)
	expected := "Movie, showing at 12pm"

	if showing != expected {
		t.Errorf("Expected showing string `%s`, found `%s`", expected, showing)
	}

	showing = movie.DisplayNextShowing(time1)

	if showing != expected {
		t.Errorf("Expected showing string `%s`, found `%s`", expected, showing)
	}
}
