// Package auth  user : benbenxiong  time : 2023-07-2023/6/7 20:20:59
package auth

import (
	"fmt"
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

	// 请求数据转结构体
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println("IsPhoneExist ShouldBindJSON err:", err.Error())
		return
	}

	// 验证手机号
	errs := requests.ValidateSignupPhoneExist(&request)
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": errs,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExits(request.Phone),
	})
}

func (s *SignupController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println("IsEmailExist  ShouldBindJSON err:", err.Error())
		return
	}

	errs := requests.ValidateSignupEmailExist(&request)
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": errs,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exits": user.IsEmailExist(request.Email),
	})

}
