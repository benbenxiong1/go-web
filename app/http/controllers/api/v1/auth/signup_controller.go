// Package auth  user : benbenxiong  time : 2023-07-2023/6/7 20:20:59
package auth

import (
	"github.com/gin-gonic/gin"
	v1 "go-web/app/http/controllers/api/v1"
	"go-web/app/model/user"
	"go-web/app/requests"
	"net/http"
)

type SignupController struct {
	v1.BaseApiController
}

func (s *SignupController) IsPhoneExist(c *gin.Context) {

	request := requests.SignupPhoneExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExits(request.Phone),
	})
}

func (s *SignupController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist); !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exits": user.IsEmailExist(request.Email),
	})

}
