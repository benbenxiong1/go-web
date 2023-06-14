// Package user  user : benbenxiong  time : 2023-07-2023/6/7 20:20:48
package user

import "go-web/pkg/database"

// IsEmailExist 判断email是否存在
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(&User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExits 判断手机号是否存在
func IsPhoneExits(phone string) bool {
	var count int64
	database.DB.Model(&User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

func GetByMulti(loginId string) (userModel User) {
	database.DB.Where("phone = ?", loginId).
		Or("email = ?", loginId).
		Or("name = ?", loginId).
		First(&userModel)
	return
}

func GetByPhone(phone string) (userModel User) {
	database.DB.Where("phone = ?", phone).
		First(&userModel)
	return
}
