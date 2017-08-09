package main

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestCLI(t *testing.T) {
	out, err := exec.Command(`./movie-recommendations`, `--genre=animation`, `--showing=10:00`).Output()

	fmt.Println(string(out))
	if err != nil {
		fmt.Println(err)
		t.Error("Error when trying to run the movie recommendation command")
	}
	expected :=
		`Zootopia, showing at 7pm
Shaun The Sheep, showing at 7pm
`
	if string(out) != expected {
		t.Errorf("Expected returned string `%s`, found `%s`", expected, string(out))
	}
}
