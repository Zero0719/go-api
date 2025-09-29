package redis

import (
	"context"
	"go-api/internal/config"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func Init() {
	config := config.Get().Redis

	Redis = redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + strconv.Itoa(config.Port),
		Password: config.Password,
		DB:       config.Database,
		PoolSize: config.PoolSize,
	})
	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("连接redis失败: %v", err)
	}
}