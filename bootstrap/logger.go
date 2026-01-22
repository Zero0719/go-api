package bootstrap

import (
	"github.com/Zero0719/go-api/pkg/config"
	"github.com/Zero0719/go-api/pkg/logger"
)

func SetupLogger() {
	logger.InitLogger(
		config.Get[string]("log.filename"),
		config.Get[int]("log.max_size"),
		config.Get[int]("log.max_backup"),
		config.Get[int]("log.max_age"),
		config.Get[bool]("log.compress"),
		config.Get[string]("log.type"),
		config.Get[string]("log.level"),
	)
}