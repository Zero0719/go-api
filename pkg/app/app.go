package app

import (
	"time"

	"github.com/Zero0719/go-api/pkg/config"
)

func IsLocal() bool {
	return config.Get[string]("app.env") == "local"
}

func IsProduction() bool {
	return config.Get[string]("app.env") == "production"
}

func IsTesting() bool {
	return config.Get[string]("app.env") == "testing"
}

func TimenowInTimezone() time.Time {
	location, _ := time.LoadLocation(config.Get[string]("app.timezone"))
	return time.Now().In(location)
}