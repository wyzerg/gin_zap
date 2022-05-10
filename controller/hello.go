package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func HelloHandler(c *gin.Context) {
	zap.L().Info("req hello")
	zap.L().Error("测试错误日志")
	c.JSON(200, gin.H{"code": 0, "msg": "你好"})
}
