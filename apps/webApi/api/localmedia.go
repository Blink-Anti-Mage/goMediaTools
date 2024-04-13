package api

import (
	"fmt"
	"goMediatools/apps/webApi/service"
	"goMediatools/datacache"
	"goMediatools/internal/ginexpand/restful"
	"goMediatools/model"
	"io/ioutil"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetLocalMedia(c *gin.Context) {
	service.Getlocal()
	restful.OkWithData(datacache.LocalCache, c)
	return
}

func GetLocalNfo(c *gin.Context) {
	var movie model.PathReq
	if err := c.ShouldBindJSON(&movie); err != nil {
		restful.FailCodeM(400, "Invalid request body", c)
		return
	}

	dir, filen := filepath.Split(movie.Path)
	fileame := filen[:len(filen)-len(filepath.Ext(filen))]
	nfoname := fileame + ".nfo"

	fmt.Println(filepath.Join(dir, nfoname))
	data, err := ioutil.ReadFile(filepath.Join(dir, nfoname))
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
