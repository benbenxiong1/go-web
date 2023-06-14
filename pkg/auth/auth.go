package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-web/app/model/user"
	"go-web/pkg/logger"
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

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}
	return userModel
}

// CurrentUid 获取当前登录用户ID
func CurrentUid(c *gin.Context) string {
	return c.GetString("current_user_id")
}