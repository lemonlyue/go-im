package router

import (
	"gin-skeleton/app/controller"
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

	// Index
	app.GET("/index/index", new(controller.IndexController).Index)
	// Add
	app.POST("/index/add", new(controller.IndexController).Add)
}
