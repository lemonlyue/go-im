package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(app *gin.Engine) {
	// 404
	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"message": "Not Found",
		})
	})
	// 测试路由
	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"message": "pong",
		})
	})
}
