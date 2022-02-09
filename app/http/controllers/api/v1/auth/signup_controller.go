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
func (s *SignupController) IsPhoneExist(c *gin.Context) {

	request := requests.SignupPhoneExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}
	//  检查数据库并返回响应
	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

// IsEmailExist 检测邮箱是否被注册
func (s *SignupController) IsEmailExist(c *gin.Context) {

	request := requests.SignupPhoneEmailRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist); !ok {
		return
	}
	//  检查数据库并返回响应
	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Email),
	})
}

// SignupUsingPhone 使用手机和验证码进行注册
func (s *SignupController) SignupUsingPhone(c *gin.Context) {

	// 1. 验证表单
	request := requests.SignupUsingPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}

	// 2. 验证成功，创建数据
	_user := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}
	_user.Create()

	if _user.ID > 0 {
		response.CreatedJSON(c, gin.H{
			"data": _user,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}

func (s SignupController) SignupUsingEmail(c *gin.Context) {
	//验证表单
	request := requests.SignupUsingEmailRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingEmail); !ok {
		return
	}

	// 2. 验证成功，创建数据
	userModel := user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	userModel.Create()

	if userModel.ID > 0 {
		response.CreatedJSON(c, gin.H{
			"data": userModel,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}

}
