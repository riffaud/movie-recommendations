package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	out, err := exec.Command(`./movie-recommendations`, `--genre=animation`, `--showing=00:00`).Output()
	if err != nil {
		t.Error("Error when trying to run the movie recommendation command")
	}

	returned := string(out)

	if !strings.Contains(returned, "Zootopia, showing at") || !strings.Contains(returned, "Shaun The Sheep, showing at") {
		t.Errorf("Expected returned string contains `%s` and `%s`, found `%s`", "Zootopia, showing at", "Shaun The Sheep, showing at", string(out))
	}
}

func TestCLIWrongTime(t *testing.T) {
	out, _ := exec.Command(`./movie-recommendations`, `--showing=dfdsfs`).Output()

	if !strings.Contains(string(out), "Error Parsing date, should use format: 15:04") {
		t.Error("Missing expected error when incorrect date format")
	}
}
