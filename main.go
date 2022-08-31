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

	// init db
	bootstrap.SetupDB()

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