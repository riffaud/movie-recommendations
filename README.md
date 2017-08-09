# Movie Recommendations CLI tool [![build status](https://travis-ci.org/riffaudo/movie-recommendations.svg?branch=master)](https://travis-ci.org/riffaudo/movie-recommendations)

Golang CLI tool to search movie recommendations based on time ang genre.

## Prerequisites

The package depends on [Golang](https://golang.org/doc/install).

## Installation/Usage

To install, checkout this repository, and run at the root of the project:

    go build

To use, run:

    ./movie-recommendations --genre="animation" --showing="00:00"

Both parameters are optional.

## Running the tests

The acceptance tests require that you build first, then simply run:

    go test -v ./...
