// Package user  user : benbenxiong  time : 2023-07-2023/6/7 20:20:48
package user

import (
	"github.com/gin-gonic/gin"
	"go-web/pkg/app"
	"go-web/pkg/database"
	"go-web/pkg/paginator"
)

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
	database.DB.Where("phone = ?", phone).First(&userModel)
	return
}

func GetByEmail(email string) (userModel User) {
	database.DB.Where("email = ?", email).First(&userModel)
	return
}

func Get(userId string) (userModel User) {
	database.DB.Where("id = ?", userId).First(&userModel)
	return
}

func All() (users []User) {
	database.DB.Find(&users)
	return
}

func Paginate(c *gin.Context, perPage int) (users []User, paging paginator.Paging) {
	paging = paginator.Paginate(c, database.DB.Model(&User{}), &users, app.V1Url(database.TableName(&User{})), perPage)

	return
}
