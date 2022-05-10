package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"gin_zap/controller"
	"gin_zap/middleware"
	"gin_zap/setting"
)

func Setup() *gin.Engine {
	gin.SetMode(setting.Conf.Mode)

	r := gin.New()
	// 注册两个自定义的中间件
	r.Use(middleware.GinLogger(zap.L()), middleware.GinRecovery(zap.L(), false))

	r.GET("/hello", controller.HelloHandler)
	r.GET("/ping", controller.PingHandler)
	r.GET("/login", controller.LoginHandler)

	return r
}
