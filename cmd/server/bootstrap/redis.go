package bootstrap

import "go-api/internal/redis"

func initRedis() {
	redis.Init()
}