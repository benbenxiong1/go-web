// Package auth  user : benbenxiong  time : 2023-07-2023/6/7 20:20:59
package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "go-web/app/http/controllers/api/v1"
	"go-web/app/model/user"
	"net/http"
)

type SignupController struct {
	v1.BaseApiController
}

func (s *SignupController) IsPhoneExist(c *gin.Context) {
	type PhoneExistRequest struct {
		Phone string `json:"phone"`
	}
	request := PhoneExistRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println("IsPhoneExist ShouldBindJSON err:", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExits(request.Phone),
	})
}
