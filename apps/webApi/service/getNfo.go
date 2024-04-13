package service

import (
	"encoding/xml"
	"goMediatools/model"
)

func GetNFO(data []byte) (*model.NfoMovie, error) {
	var movie model.NfoMovie
	err := xml.Unmarshal(data, &movie)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}
