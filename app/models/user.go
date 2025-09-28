package models

import (
	"go-api/app/utils"

	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Username string `gorm:"type:varchar(64);not null;default:'';unique"`
	Password string `gorm:"type:char(32);not null;default:''"`
}

func (m *User) BeforeCreate(tx *gorm.DB) (err error) {
	m.Password = utils.MD5(m.Password + utils.Config.GetString("app.salt"))
	return
}

func (m *User) GetByUsername(username string) (User, error) {
	var user User
	if err := DB.Where("username = ?", username).Limit(1).Find(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (m *User) Create() error {
	if err := DB.Create(m).Error; err != nil {
		return err
	}
	return nil
}