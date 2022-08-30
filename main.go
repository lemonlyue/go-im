package main

import (
	"flag"
	"gin-skeleton/bootstrap"
	"gin-skeleton/provider"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func main()  {
	// init config
	provider.InitConfig()
	provider.InitGormDB()

	addr := flag.String("addr", ":" + viper.GetString("Server.Http.Port"), "Address to listen and serve")
	app := gin.Default()
	bootstrap.SetupRoute(app)
	err := app.Run(*addr)
	// Printf Error
	if err != nil {
		log.Fatal(err.Error())
	}
}