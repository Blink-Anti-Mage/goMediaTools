package api

import (
	"goMediatools/apps/webApi/service"
	"goMediatools/datacache"
	"goMediatools/internal/ginexpand/restful"
	"goMediatools/model"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetLocalMedia(c *gin.Context) {
	service.Getlocal()
	restful.OkWithData(datacache.LocalCache, c)
	return
}

func GetFileMedia(c *gin.Context) {
	var path model.PathReq
	if err := c.ShouldBindJSON(&path); err != nil {
		restful.FailCodeM(400, "Req error", c)
		return
	}
	tree, err := service.BuildTree(path.Path)
	if err != nil {
		restful.FailCodeM(400, "Invalid request body", c)
		return
	}

	restful.OkWithData(tree, c)
	return
}

func GetLocalNfo(c *gin.Context) {
	var movie model.PathReq
	if err := c.ShouldBindJSON(&movie); err != nil {
		restful.FailCodeM(400, "Invalid request body", c)
		return
	}

	data, err := os.ReadFile(movie.Path)
	if err != nil {
		restful.FailWithMessage("nfo ioread error", c)
		return
	}
	jsondata, err := service.GetNFO(data)
	if err != nil {
		restful.FailWithMessage("nfo ioread error", c)
		return
	}
	restful.OkWithData(jsondata, c)
	return
}

func RunMovedir(c *gin.Context) {
	var path model.PathReq
	if err := c.ShouldBindJSON(&path); err != nil {
		restful.FailCodeM(400, "Invalid request body", c)
		return
	}
	err := service.MoveDir(path.Path)
	if err != nil {
		restful.FailWithMessage("RunMovedir  error", c)
		return
	}

	restful.Ok(c)
	return
}

func Renameworker(c *gin.Context) {
	var path model.PathReq
	if err := c.ShouldBindJSON(&path); err != nil {
		restful.FailCodeM(400, "Invalid request body", c)
		return
	}
	err := service.MoveDir2(path.Path)
	if err != nil {
		restful.FailWithMessage("RunMovedir  error", c)
		return
	}

	restful.Ok(c)
	return
}

func PutNfo(c *gin.Context) {
	var movie model.NfoReq
	if err := c.ShouldBindJSON(&movie); err != nil {
		restful.FailCodeM(400, err.Error(), c)
		return
	}

	xmldata, err := service.SetNFO(movie)
	if err != nil {
		restful.FailWithMessage("nfo ioread error", c)
		return
	}

	if err := os.WriteFile(movie.Path, xmldata, 0644); err != nil {
		restful.FailWithMessage("nfo WriteFile error", c)
		return
	}

	restful.OkWithMessage("", c)
	return
}

func GetTmdbNfo(c *gin.Context) {
	var req model.TmdbNfoReq
	if err := c.ShouldBindJSON(&req); err != nil {
		restful.FailCodeM(400, err.Error(), c)
		return
	}

	data, err := service.GetMovieDetails(req.Tmdbid, "")
	if err != nil {
		restful.FailCodeM(400, err.Error(), c)
		return
	}

	var movie model.NfoMovie
	movie.Title = data.Title
	movie.Plot = data.Overview
	movie.Tmdbid = req.Tmdbid
	movie.Year = "0000"
	t, err := time.Parse("2006-01-02", data.ReleaseDate)
	if err != nil {
	} else {
		movie.Year = strconv.Itoa(t.Year())
	}

	xmldata, err := service.SetNFOs(movie)
	if err != nil {
		restful.FailWithMessage("nfo ioread error", c)
		return
	}

	dir, filen := filepath.Split(req.Path)
	fileame := filen[:len(filen)-len(filepath.Ext(filen))]
	nfoname := fileame + ".nfo"

	// 将XML数据写入文件
	if err := os.WriteFile(filepath.Join(dir, nfoname), xmldata, 0644); err != nil {
		restful.FailWithMessage("nfo WriteFile error", c)
		return
	}

	//  保存海报
	if data.PosterPath != "" {
		posterUrl := "https://image.tmdb.org/t/p/w500/" + data.PosterPath
		ext := filepath.Ext(data.PosterPath)
		err = service.DownloadPoster(posterUrl, filepath.Join(dir, fileame+"-poster"+ext))

		if err != nil {
			restful.FailWithMessage("poster WriteFile error", c)
			return
		}
	}

	restful.OkWithMessage("", c)
	return
}
