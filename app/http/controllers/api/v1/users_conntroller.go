package v1

import (
	"github.com/gin-gonic/gin"
	"go-web/app/model/user"
	"go-web/app/requests"
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
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := user.Paginate(c, 10)
	response.Json(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}
