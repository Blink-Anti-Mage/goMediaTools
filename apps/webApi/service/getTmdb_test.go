package service

import (
	"testing"
)

func TestGetTmdbMovieInfo(t *testing.T) {
	movieName := "长津湖"
	GetMovieInfo(movieName, "")

}

func TestGetTmdbMovieDetails(t *testing.T) {
	movieid := 779029
	GetMovieDetails(movieid, "")

}

func TestGetTmdbMovieCredits(t *testing.T) {
	movieid := 779029
	GetMovieCredits(movieid, "")

}
