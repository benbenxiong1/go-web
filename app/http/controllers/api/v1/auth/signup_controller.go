package auth

import (
	"github.com/gin-gonic/gin"
	v1 "go-web/app/http/controllers/api/v1"
	"go-web/app/models/user"
	"go-web/app/requests"
	"go-web/pkg/response"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (s *SignupController) IsPhoneExist(c *gin.Context)  {

	request := requests.SignupPhoneExistRequest{}

	if ok := requests.Validate(c,&request,requests.ValidateSignupPhoneExist);!ok{
		return
	}
	//  检查数据库并返回响应
	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

// IsEmailExist 检测邮箱是否被注册
func (s *SignupController) IsEmailExist(c *gin.Context)  {

	request := requests.SignupPhoneEmailRequest{}

	if ok := requests.Validate(c,&request,requests.ValidateSignupEmailExist);!ok{
		return
	}
	//  检查数据库并返回响应
	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Email),
	})
}