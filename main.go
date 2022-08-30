package main

import (
	"flag"
	"gin-skeleton/router"
	"gin-skeleton/provider"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main()  {
	// init config
	provider.InitConfig()

	addr := flag.String("addr", ":" + viper.GetString("Server.Http.Port"), "Address to listen and serve")
	app := gin.Default()
	router.Router(app)
	app.Run(*addr)
}