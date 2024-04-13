package service

import (
	"encoding/json"
	"fmt"
	"goMediatools/datacache"
	"goMediatools/internal/httpclient"
	"goMediatools/model"
)

var headermap = map[string]string{
	"accept":        "application/json",
	"Authorization": "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiIyOTU1YmUzMTU0ZWZjNzY3ODA4ODQ0YTEzZmU5MTU5NSIsInN1YiI6IjY1ZmE3NzM2NzcwNzAwMDE0OTA1Y2JiZiIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.rYxRKeWRLpsLTwoykSRzdA8u5SwLZ7hOZ0YHOD1YeHE",
}

func GetMovieInfo(movieName string, language string) {
	if language == "" {
		language = datacache.CN
	}
	url := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?query=%s&include_adult=false&language=%s&page=1&year=2021", movieName, language)
	fmt.Println(url)
	resp, err := httpclient.GetV2(url, headermap, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	var data model.SearchMovie
	if err := json.Unmarshal(resp, &data); err != nil {
		fmt.Println(err.Error())
	}
	for _, v := range data.Results {
		datacache.TMDBCache[v.ID] = v
	}
	temp := datacache.LocalCache[movieName]
	temp.Tmdbid = data.Results[0].ID
	datacache.LocalCache[movieName] = temp

	fmt.Println(datacache.LocalCache)
}

func GetMovieDetails(movieId int, language string) {
	if language == "" {
		language = datacache.CN
	}
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%d?language=%s", movieId, language)
	fmt.Println(url)
	resp, err := httpclient.GetV2(url, headermap, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(resp))
}

func GetMovieCredits(movieId int, language string) {
	if language == "" {
		language = datacache.CN
	}
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%d/credits?language=%s", movieId, language)
	fmt.Println(url)
	resp, err := httpclient.GetV2(url, headermap, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(resp))
}