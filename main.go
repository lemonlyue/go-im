package main

import (
	"flag"
	"gin-skeleton/bootstrap"
	bootstrapConfig "gin-skeleton/config"
	"gin-skeleton/pkg/config"
	"github.com/gin-gonic/gin"
	"log"
)

func init()  {
	bootstrapConfig.Initialize()
}

func main()  {
	// Init Config
	var env string
	flag.StringVar(&env, "env", "", "load .env file")
	flag.Parse()
	config.InitConfig(env)

	// init logger
	bootstrap.SetupLogger()

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	// init db
	bootstrap.SetupDB()

	// init redis
	bootstrap.SetupRedis()

	addr := flag.String("addr", ":" + config.Get("app.port"), "Address to listen and serve")
	app := gin.Default()
	// init router
	bootstrap.SetupRoute(app)
	err := app.Run(*addr)
	// Printf Error
	if err != nil {
		log.Fatal(err.Error())
	}
}