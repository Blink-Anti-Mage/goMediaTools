package router

import (
	"goMediatools/apps/webApi/api"

	"github.com/gin-gonic/gin"
)

var r *gin.RouterGroup

func InitRouter(g *gin.Engine) {
	r = g.Group("/api/v1")
	api.GinEngine = g
	{
		r.GET("/ping", api.Ping)
		r.POST("/StaticFS", api.StaticFS)

		//r.GET("/getLocalMedia", api.GetLocalMedia)

		r.POST("/getFileMedia", api.GetFileMedia)
		r.POST("/getLocalNfo", api.GetLocalNfo)
		r.POST("/putLocalNfo", api.PutNfo)

		r.POST("/getTmdbNfo", api.GetTmdbNfo)

		r.POST("/runMovedir", api.RunMovedir)
		r.POST("/renameworker", api.Renameworker)
	}

}
