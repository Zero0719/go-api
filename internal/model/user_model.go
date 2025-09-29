package model

import (
	"go-api/internal/config"
	"go-api/pkg/md5"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(64);not null;default:'';unique"`
	Password string `gorm:"type:char(32);not null;default:''"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = md5.MD5(u.Password + config.Get().App.Salt)
	return
  }
