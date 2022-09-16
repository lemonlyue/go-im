package app

import (
	"gin-skeleton/pkg/config"
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
