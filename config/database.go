package config

import "gin-skeleton/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{} {
			// CONNECTION
			"connection": config.Env("DB_CONNECTION", "127.0.0.1"),
			"mysql": map[string]interface{}{
				// DATABASE CONFIG
				"host":     config.Env("DB_HOST", "127.0.0.1"),
				"port":     config.Env("DB_PORT", "3306"),
				"database": config.Env("DB_DATABASE", "gin-skeleton"),
				"username": config.Env("DB_USERNAME", "root"),
				"password": config.Env("DB_PASSWORD", "root"),
				"charset":  "utf8mb4",

				// CONNECTION POOL
				"max_idle_connections": config.Env("DB_MAX_IDLE_CONNECTIONS", 100),
				"max_open_connections": config.Env("DB_MAX_OPEN_CONNECTIONS", 25),
				"max_life_seconds":     config.Env("DB_MAX_LIFE_SECONDS", 5*60),
			},
		}
	})
}
