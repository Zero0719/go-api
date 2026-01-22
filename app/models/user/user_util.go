package user

import "github.com/Zero0719/go-api/pkg/database"

func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

func Get(idstr string) (userModel User) {
	database.DB.Where("id = ?", idstr).First(&userModel)
	return
}