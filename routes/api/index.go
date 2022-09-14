package api

import (
	"github.com/gin-gonic/gin"
	Controller "micro-gin/app/controllers"
)

func Routers(e *gin.Engine) {
	c := new(Controller.IndexController)
	e.GET("/index", c.Index)
}
git