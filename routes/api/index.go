package routes

import (
	"github.com/gin-gonic/gin"
	"micro-gin/app/controllers/app"
	"micro-gin/app/controllers/common"
	"micro-gin/app/middleware"
	"micro-gin/app/services"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/auth/register", app.Register)
	router.POST("/auth/test_gen_struct", app.TestGenStruct)
	router.GET("/auth/login", app.Login)
	router.GET("/auth/info", app.Info)

	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		//authRouter.POST("/auth/info", app.Info)
		authRouter.POST("/auth/logout", app.Logout)
		authRouter.POST("/image_upload", common.ImageUpload)
	}
}
