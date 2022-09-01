package bootstrap

import (
	"gin-skeleton/pkg/app"
	"gin-skeleton/pkg/config"
	"gin-skeleton/pkg/console"
	"gin-skeleton/pkg/logger"
	"github.com/gin-gonic/gin"
)

func SetupServe() *gin.Engine {
	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	mode := gin.ReleaseMode
	// 本地环境及测试环境使用 debug
	if app.IsLocal() || app.IsTest() {
		mode = gin.DebugMode
	}
	if app.IsUnitTest() {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	// 初始化路由绑定
	router := SetupRoute()

	// 运行服务器
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		logger.ErrorString("CMD", "serve", err.Error())
		console.Exit("Unable to start server, error:" + err.Error())
	}

	return router
}
