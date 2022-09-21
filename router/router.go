package router

import (
	"gin-skeleton/app/controller"
	"gin-skeleton/app/ws"
	"gin-skeleton/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(app *gin.Engine) {
	// 测试路由
	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "pong",
		})
	})

	// Index
	app.GET("/index/index", new(controller.IndexController).Index)
	// Add
	app.POST("/index/add", new(controller.IndexController).Add)

	// user module
	userGroup := app.Group("/user")
	{
		// login
		userGroup.POST("/login", new(controller.UserController).Login)
		// register
		userGroup.POST("/register", new(controller.UserController).Register)
		// get user info
		userGroup.GET("/getUserInfo", middlewares.AuthJWT(), new(controller.UserController).GetUserInfo)
	}

	// ws
	group := app.Group("/ws")
	{
		group.GET("", new(ws.WsController).WsClient)
	}
}
