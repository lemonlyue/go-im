package app

import (
	"gin-skeleton/pkg/config"
	"time"
)

var IsUnitTestEnv = false

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}

func IsTest() bool {
	return config.Get("app.env") == "test"
}

func IsUnitTest() bool {
	return IsUnitTestEnv == true
}

// TimenowInTimezone 获取当前时间，支持时区
func TimenowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.GetString("app.timezone"))
	return time.Now().In(chinaTimezone)
}
