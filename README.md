# Movie Recommendations CLI tool [![build status](https://travis-ci.org/riffaudo/movie-recommendations.svg?branch=master)](https://travis-ci.org/riffaudo/movie-recommendations)

Golang CLI tool to search movie recommendations based on time ang genre.

## Prerequisites

The package depends on [Golang](https://golang.org/doc/install).

## Install and Run

To install, checkout this repository, and run at the root of the project:

    go build

To use, run:

    ./movie-recommendations --genre="animation" --showing="00:00"

Both parameters are optional.

## Use in your code

You can use this library directly in your code. There is only one type of storage at the moment, that is taking an io.Reader as an argument.

    import "github.com/riffaudo/movie-recommendations/movie"

    storage, _ := movie.StorageFromReader(io.Reader)
    moviesData := storage.Load(movie.SearchParams{Genre: "Animation"})

## Running the tests

The acceptance tests require that you build first:

    go build

 then run:

    go test -v ./...
