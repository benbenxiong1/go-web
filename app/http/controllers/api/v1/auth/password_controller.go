// Package auth  user : benbenxiong  time : 2023-14-2023/6/14 20:20:18
package auth

import (
	"github.com/gin-gonic/gin"
	v1 "go-web/app/http/controllers/api/v1"
	"go-web/app/model/user"
	"go-web/app/requests"
	"go-web/pkg/response"
)

type PasswordController struct {
	v1.BaseApiController
}

func (p *PasswordController) ResetByPhone(c *gin.Context) {
	request := requests.ResetByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByPhone); !ok {
		return
	}

	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()

		response.Success(c)
	}

}

func (p *PasswordController) ResetByEmail(c *gin.Context) {
	request := requests.ResetByEmailRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByEmail); !ok {
		return
	}

	userModel := user.GetByEmail(request.Email)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()

		response.Success(c)
	}
}
