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

func SetNFO(data model.NfoReq) (xmldata []byte, err error) {
	var movie model.NfoMovie
	movie.Title = data.Title
	movie.Plot = data.Plot

	xmldata, err = xml.MarshalIndent(movie, "", " ")
	if err != nil {
		return nil, err
	}

	return xmldata, nil
}

func SetNFOs(data model.NfoMovie) (xmldata []byte, err error) {
	xmldata, err = xml.MarshalIndent(data, "", " ")
	if err != nil {
		return nil, err
	}

	return xmldata, nil
}
