package bootstrap

import (
	"gin-skeleton/middlewares"
	"github.com/gin-gonic/gin"
	routers "gin-skeleton/router"
	"net/http"
)

func SetupRoute() *gin.Engine {
	// gin 实例
	router := gin.New()

	// Register Global Middleware
	registerGlobalMiddleWare(router)

	// Register Router
	routers.RegisterRouter(router)

	// Handle 404 Not Found
	setup404Handler(router)

	return router
}

func registerGlobalMiddleWare(router *gin.Engine)  {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
	)
}

func setup404Handler(router *gin.Engine)  {
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"message": "Not Found",
		})
	})
}
