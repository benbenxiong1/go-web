// Package user  user : benbenxiong  time : 2023-07-2023/6/7 20:20:44
package user

import (
	"go-web/app/model"
	"go-web/pkg/database"
)

type User struct {
	model.BaseModel
	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	model.CommonTimestampsField
}

func (u *User) Create() {
	database.DB.Create(&u)
}
