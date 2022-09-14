package bootstrap

import (
	"github.com/gin-gonic/gin"
	"micro-gin/routes/api"
)

var Router *gin.Engine

type Option func(*gin.Engine)

var options = []Option{}

func initRouter() *gin.Engine {
	Router = gin.New()
	Include(api.Routers)
	//r.Use(Global.CommonMiddleWare(),Global.Myloger())
	for _, opt := range options {
		opt(Router)
	}
	return Router
}

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}
