package router

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(app *gin.Engine) {
	// 测试路由
	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"message": "pong",
		})
	})
}
