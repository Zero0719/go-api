package bootstrap

import (
	"fmt"
	"github.com/Zero0719/go-api/app/models/user"
	"time"

	"github.com/Zero0719/go-api/pkg/config"
	"github.com/Zero0719/go-api/pkg/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger "github.com/Zero0719/go-api/pkg/logger"
)

func SetupDB() {
	var dbConfig gorm.Dialector

	switch config.Get[string]("database.connection") {
	case "mysql":
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%s&parseTime=True&loc=Local&multiStatements=true", config.Get[string]("database.mysql.username"), config.Get[string]("database.mysql.password"), config.Get[string]("database.mysql.host"), config.Get[string]("database.mysql.port"), config.Get[string]("database.mysql.database"), config.Get[string]("database.mysql.charset"))
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	default:
		panic("Unsupported database driver: " + config.Get[string]("database.connection"))
	}

	database.Connect(dbConfig, logger.NewGormLogger())

	database.SqlDB.SetMaxIdleConns(config.Get[int]("database.mysql.max_idle_connections"))
	database.SqlDB.SetMaxOpenConns(config.Get[int]("database.mysql.max_open_connections"))
	database.SqlDB.SetConnMaxLifetime(time.Duration(config.Get[int]("database.mysql.max_life_seconds")) * time.Second)

	database.DB.AutoMigrate(&user.User{})
}