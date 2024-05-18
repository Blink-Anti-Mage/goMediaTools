package service

import (
	"encoding/json"
	"fmt"
	"goMediatools/model"
	"testing"
)

func TestGetTmdbMovieInfo(t *testing.T) {
	movieName := "长津湖"
	GetMovieInfo(movieName, "")

}

func TestGetTmdbMovieDetails(t *testing.T) {
	movieid := "779029"
	GetMovieDetails(movieid, "")

}

func TestGetTmdbMovieCredits(t *testing.T) {
	movieid := 779029
	GetMovieCredits(movieid, "")

}

func TestJson(t *testing.T) {
	movieName := `{"adult":false,"backdrop_path":"/IRN1JuNwr1VKp88Dscgja2uR8H.jpg","belongs_to_collection":null,"budget":80000000,"genres":[{"id":10752,"name":"战争"},{"id":36,"name":"历史"},{"id":18,"name":"剧情"},{"id":28,"name":"动作"}],"homepage":"","id":508935,"imdb_id":"tt7294150","origin_country":["CN"],"original_language":"zh","original_title":"八佰","overview":"故事原型为一九三七年发生于上海的四行仓库保卫战，此战为淞沪会战最后一役，故事围绕“八百孤军血战四行仓库”展开。1937年淞沪会战末期，中日双方激战已持续三个月，上海濒临沦陷。第88师262旅524团团附谢晋元（杜淳饰）率420余人，孤军坚守最后的防线，留守上海四行仓库。与租界一河之隔，造就了罕见的被围观的战争。为壮声势，实际人数四百人而对外号称八百人 。“八百壮士”奉命留守上海闸北，在苏州河畔的四行仓库鏖战四天，直至10月30日才获令撤往英租界。","popularity":25.625,"poster_path":"/kKUhjtExjI9caXU2xGJXZOD5Bnf.jpg","production_companies":[{"id":112270,"logo_path":null,"name":"Beijing Diqi Yinxiang Entertainment","origin_country":""},{"id":3393,"logo_path":null,"name":"Huayi Brothers Pictures","origin_country":"CN"},{"id":17818,"logo_path":"/a0dRBd4hPab4KjmbiIJpNwRfIVD.png","name":"Beijing Enlight Pictures","origin_country":"CN"},{"id":81620,"logo_path":"/gNp4dfuBOXmVWdGKb63NfbFNbFi.png","name":"Tencent Pictures","origin_country":"CN"},{"id":2806,"logo_path":"/vxOhCbpsRBh10m6LZ3HyImTYpPY.png","name":"South Australian Film Corporation","origin_country":"AU"},{"id":25660,"logo_path":null,"name":"Huaxia Film Distribution","origin_country":"CN"}],"production_countries":[{"iso_3166_1":"AU","name":"Australia"},{"iso_3166_1":"CN","name":"China"}],"release_date":"2020-08-14","revenue":461421559,"runtime":147,"spoken_languages":[{"english_name":"English","iso_639_1":"en","name":"English"},{"english_name":"Mandarin","iso_639_1":"zh","name":"普通话"},{"english_name":"Japanese","iso_639_1":"ja","name":"日本語"}],"status":"Released","tagline":"","title":"八佰","video":false,"vote_average":6.973,"vote_count":240}`
	var data model.MovieDetails
	if err := json.Unmarshal([]byte(movieName), &data); err != nil {
		fmt.Println(data)
	}
	fmt.Println(data.ID, data.Overview)
}
