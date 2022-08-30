package provider

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"strings"
	"gorm.io/driver/mysql"
	"time"
)

// Init GormDb
func InitGormDB()  {
	gormMysql("default")
}

func gormMysql(connection string) *gorm.DB  {
	connection = strings.ToUpper(connection)
	host := viper.GetString("Gorm." + connection + ".Host")
	port := viper.GetInt("Gorm." + connection + ".Port")
	database := viper.GetString("Gorm." + connection + ".Database")
	username := viper.GetString("Gorm." + connection + ".Username")
	password := viper.GetString("Gorm." + connection + ".Password")
	charset := viper.GetString("Gorm." + connection + ".Charset")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local", username, password, host, port, database, charset)
	mysqlConfig := mysql.Config{
		DSN: dsn,// DSN data source name
	}
	
	db, err := gorm.Open(mysql.New(mysqlConfig))
	if err != nil {
		log.Println("Gorm connection error: ", err, connection)
		return nil
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
