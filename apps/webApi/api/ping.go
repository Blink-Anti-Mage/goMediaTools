package api

import (
	"goMediatools/internal/ginexpand/restful"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {

	restful.OkWithMessage("ping success", c)
}
