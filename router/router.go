package router

import (
	"github.com/gin-gonic/gin"
)

func Router(app *gin.Engine)  {
	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
