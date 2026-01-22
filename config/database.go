package config

import "github.com/Zero0719/go-api/pkg/config"

func init() {
	config.Add("database", func() map[string]interface{} {
		return map[string]interface{}{
			"connection": config.Env("DB_CONNECTION", "mysql"),
			
			"mysql": map[string]interface{}{
				"host": config.Env("DB_HOST", "127.0.0.1"),
				"port": config.Env("DB_PORT", 3306),
				"database": config.Env("DB_DATABASE", "go-api"),
				"username": config.Env("DB_USERNAME", "root"),
				"password": config.Env("DB_PASSWORD", ""),
				"charset": config.Env("DB_CHARSET", "utf8mb4"),

				"max_idle_connections": config.Env("DB_MAX_IDLE_CONNECTIONS", 100),
				"max_open_connections": config.Env("DB_MAX_OPEN_CONNECTIONS", 25),
				"max_life_seconds": config.Env("DB_MAX_LIFE_SECONDS", 300),
			},
		}
	})
}