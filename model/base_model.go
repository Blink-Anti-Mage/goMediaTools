package model

type NameReq struct {
	Name string `json:"name"`
}

type PathReq struct {
	Path string `json:"path"`
}

type NfoReq struct {
	Plot  string `json:"plot"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

type TmdbNfoReq struct {
	Tmdbid string `json:"tmdbid"`
	Path   string `json:"path"`
}
