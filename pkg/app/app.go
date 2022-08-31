package app

import "gin-skeleton/pkg/config"

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}

func IsTest() bool {
	return config.Get("app.env") == "test"
}
