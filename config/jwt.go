package config

import "github.com/Zero0719/go-api/pkg/config"

func init() {
	config.Add("jwt", func() map[string]interface{} {
		return map[string]interface{}{
			"max_refresh_time": config.Env("JWT_MAX_REFRESH_TIME", 60),
			"expire_time": config.Env("JWT_EXPIRE_TIME", 60),
			"debug_expire_time": config.Env("JWT_DEBUG_EXPIRE_TIME", 120),
		}
	})
}