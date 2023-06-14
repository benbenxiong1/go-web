package auth

import (
	"errors"
	"go-web/app/model/user"
)

// Attempt 尝试登录
func Attempt(loginId, password string) (user.User, error) {
	userModel := user.GetByMulti(loginId)
	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}

	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}

	return userModel, nil
}

// LoginByPhone 手机号登录
func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("手机号未注册")
	}
	return userModel, nil
}
