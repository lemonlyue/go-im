package provider

import (
	//"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

// init config
func InitConfig()  {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	// read yaml config
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Read config failed: ", err)
	}

	// listen config change
	//viper.WatchConfig()
	//viper.OnConfigChange(func(e fsnotify.Event) {
	//	log.Println("Config file changed: ", e.Name)
	//})
}
