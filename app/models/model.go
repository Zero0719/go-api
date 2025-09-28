package models

import (
	"fmt"
	"go-api/app/utils"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	if utils.Config.GetString("db.type") != "mysql" {
		return
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", utils.Config.GetString("db.user"), utils.Config.GetString("db.password"), utils.Config.GetString("db.host"), utils.Config.GetInt("db.port"), utils.Config.GetString("db.database"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Logger.Error().Msgf("failed to connect database: %v", err)
		log.Fatalf("failed to connect database: %v", err)
	}
	DB = db
	if utils.Config.GetString("app.env") == "development" {
		DB.AutoMigrate(&User{})
	}
}

type BaseModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}