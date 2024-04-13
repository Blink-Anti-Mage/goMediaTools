package ginexpand

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// func RequestIdMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("X-Request-Id", uuid.New().String())
// 		c.Next()
// 	}
// }

func CORSMiddleware() gin.HandlerFunc {
	cfg := cors.Default()
	return cfg
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, UPDATE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
