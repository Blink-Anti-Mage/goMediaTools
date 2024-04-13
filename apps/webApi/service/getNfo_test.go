package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestGetNFO(t *testing.T) {
	data, err := ioutil.ReadFile(`\\nas\media\电影\57.义胆群英 (1989)DVD-REMUX\義膽群英.1989.nfo`)
	if err != nil {

	}
	movie, err := GetNFO(data)
	if err != nil {

	}
	jsonData, err := json.Marshal(movie)
	if err != nil {
		fmt.Printf("JSON marshaling failed: %s", err)
	}
	fmt.Print(string(jsonData))

}
