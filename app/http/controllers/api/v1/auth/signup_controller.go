// Package auth  user : benbenxiong  time : 2023-07-2023/6/7 20:20:59
package auth

import (
	"github.com/gin-gonic/gin"
	v1 "go-web/app/http/controllers/api/v1"
	"go-web/app/model/user"
	"go-web/app/requests"
	"go-web/pkg/jwt"
	"go-web/pkg/response"
)

type SignupController struct {
	v1.BaseApiController
}

func (s *SignupController) IsPhoneExist(c *gin.Context) {

	request := requests.SignupPhoneExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}

	response.Json(c, gin.H{
		"exist": user.IsPhoneExits(request.Phone),
	})
}

func (s *SignupController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist); !ok {
		return
	}

	response.Json(c, gin.H{
		"exits": user.IsEmailExist(request.Email),
	})

}

func (s *SignupController) SignupUsingPhone(c *gin.Context) {
	request := requests.SignupUsingPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}

	// 验证通过
	_user := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}

	_user.Create()

	if _user.ID > 0 {
		token := jwt.NewJwt().IssueToken(_user.GetStringId(), _user.Name)
		response.CreatedJSON(c, gin.H{
			"data":  _user,
			"token": token,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}

}

func (s *SignupController) SignupUsingEmail(c *gin.Context) {
	request := requests.SignupUsingEmailRequest{}

	if ok := requests.Validate(c, &request, requests.SignupUsingEmail); !ok {
		return
	}

	// 验证通过
	_user := user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	_user.Create()

	if _user.ID > 0 {
		token := jwt.NewJwt().IssueToken(_user.GetStringId(), _user.Name)
		response.Created(c, gin.H{
			"data":  _user,
			"token": token,
		})
	} else {
		response.Abort500(c, "添加失败，请重试")
	}

}
