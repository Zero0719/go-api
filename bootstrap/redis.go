package bootstrap

import (
	"fmt"

	"github.com/Zero0719/go-api/pkg/config"
	"github.com/Zero0719/go-api/pkg/redis"
)

func SetupRedis() {
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.Get[string]("redis.host"), config.Get[string]("redis.port")),
		config.Get[string]("redis.username"),
		config.Get[string]("redis.password"),
		config.Get[int]("redis.database"),
	)
}