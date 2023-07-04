package v1

import (
	"github.com/gin-gonic/gin"
	"go-web/pkg/auth"
	"go-web/pkg/response"
)

type UsersController struct {
	BaseApiController
}

// CurrenUser 当前登录用户信息
func (ctrl *UsersController) CurrenUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}
