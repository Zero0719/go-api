package utils

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func InitRedis() {
	if Config.GetString("cache.type") != "redis" {
		return
	}
	Redis = redis.NewClient(&redis.Options{
		Addr:     Config.GetString("cache.host") + ":" + Config.GetString("cache.port"),
		Password: Config.GetString("cache.password"),
		DB:       Config.GetInt("cache.database"),
		PoolSize: Config.GetInt("cache.pool_size"),
	})
	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		Logger.Fatal().Msgf("连接redis失败: %v", err)
	}
}