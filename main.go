package main

import (
	"fmt"
	"goMediatools/internal/config"
	"goMediatools/internal/ginexpand"
	"goMediatools/router"

	"github.com/gin-gonic/gin"
)

func main() {

	err := config.InitConfig("./config.json")
	if err != nil {
		fmt.Println("add config err:" + err.Error())
		return
	}

	r := gin.Default()
	r.Use(ginexpand.Cors())
	router.InitRouter(r)

	r.Run(":5025")
}
