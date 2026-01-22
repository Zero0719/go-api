package user

import "github.com/Zero0719/go-api/app/models"

type User struct {
	models.BaseModel
	
	Name     string `json:"name,omitempty"`
    Email    string `json:"-"`
    Phone    string `json:"-"`
    Password string `json:"-"`

	models.CommonTimestampsField
}