package controller

import (
	"gin_zap/service"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	// 1. 参数处理
	// 2. 业务逻辑
	service.Login()
	// 3. 返回响应
	c.JSON(200, gin.H{"code": 0, "msg": "login success"})
}
