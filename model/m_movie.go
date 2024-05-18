package model

import "encoding/xml"

type NfoMovie struct {
	XMLName       xml.Name `xml:"movie" json:"movie"`
	Plot          string   `xml:"plot" json:"plot"`
	Lockdata      string   `xml:"lockdata" json:"lockdata"`
	Dateadded     string   `xml:"dateadded" json:"dateadded"`
	Title         string   `xml:"title"  json:"title"`
	Originaltitle string   `xml:"originaltitle" json:"originaltitle"`
	Directors     []string `xml:"director" json:"director"`
	Writers       []string `xml:"writer" json:"writer"`
	Credits       []string `xml:"credits 	 json:"credits"`
	Rating        string   `xml:"rating" json:"rating"`
	Year          string   `xml:"year" json:"year"`
	MPAA          string   `xml:"mpaa" json:"mpaa"`
	Tmdbid        string   `xml:"tmdbid" json:"tmdbid"`
}

type Movie struct {
	Name      string
	Path      string
	IfNfo     bool
	IfPoster  bool
	PosterDir string
	IfPath    bool
	Tmdbid    int
}
