package router

import (
	"goMediatools/apps/webApi/api"

	"github.com/gin-gonic/gin"
)

var r *gin.RouterGroup

func InitRouter(g *gin.Engine) {
	r = g.Group("/api/v1")

	{
		r.GET("/ping", api.Ping)

		r.GET("/getLocalMedia", api.GetLocalMedia)
		r.POST("/getLocalNfo", api.GetLocalNfo)
		r.POST("/putLocalNfo", api.PutNfo)
	}

}
