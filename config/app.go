package config

import (
	"gin-skeleton/pkg/config"
)

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{} {
			// APP NAME
			"name": config.Env("APP_NAME", "gin-skeleton"),
			// ENV
			"env": config.Env("APP_ENV", "production"),
			// DEBUG
			"debug": config.Env("APP_DEBUG", false),
			// PORT
			"port": config.Env("APP_PORT", "8080"),
			// TIMEZONE
			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),
		}
	})
}