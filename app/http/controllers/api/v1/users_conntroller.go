package v1

import (
	"github.com/gin-gonic/gin"
	"go-web/app/model/user"
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

func (ctrl *UsersController) Index(c *gin.Context) {
	data := user.All()
	response.Data(c, data)
}
