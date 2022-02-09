package user

import (
	"go-web/app/models"
)

// IsEmailExist 判断email是否存在
func IsEmailExist(email string) bool {
	return models.IsExist(User{}, "email", email)
}

func IsPhoneExist(phone string) bool {
	return models.IsExist(User{}, "phone", phone)
}
