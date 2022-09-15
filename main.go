package main

import (
	"github.com/gin-gonic/gin"
	"micro-gin/bootstrap"
	"micro-gin/global"
	"net/http"
)

func main() {

	//bootstrap.Router.Run()
	bootstrap.InitializeConfig()

	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})
	r.Run(":" + global.App.Config.App.PORT)
}
