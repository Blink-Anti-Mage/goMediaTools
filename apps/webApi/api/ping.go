package api

import (
	"goMediatools/internal/ginexpand/restful"
	"goMediatools/model"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var GinEngine *gin.Engine

func Ping(c *gin.Context) {

	restful.OkWithMessage("ping success", c)
}

func StaticFS(c *gin.Context) {
	var path model.PathReq
	if err := c.ShouldBindJSON(&path); err != nil {
		restful.FailCodeM(400, "Invalid request body", c)
		return
	}
	r := filepath.Join(path.Path, "")
	GinEngine.StaticFS("/images", gin.Dir(r, true))

	restful.OkWithMessage("StaticFS ok", c)
}
