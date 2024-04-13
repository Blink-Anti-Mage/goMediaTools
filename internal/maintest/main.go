package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int
	Genre  string
	Rating float64
}

func main() {
	movies := make(map[string]Movie)
	movies["Inception"] = Movie{"Inception", 2010, "Sci-Fi", 8.8}
	movies["Interstellar"] = Movie{"Interstellar", 2014, "Sci-Fi", 8.6}

	jsonData, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Printf("%s\n", jsonData)
}
