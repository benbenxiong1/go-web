package auth

import (
	"github.com/gin-gonic/gin"
	v1 "go-web/app/http/controllers/api/v1"
	"go-web/app/requests"
	"go-web/pkg/auth"
	"go-web/pkg/jwt"
	"go-web/pkg/response"
)

type LoginController struct {
	v1.BaseApiController
}

func (l *LoginController) LoginByPhone(c *gin.Context) {
	// 验证表单
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}

	// 尝试登录
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		response.Error(c, err, "账号不存在")
	} else {
		token := jwt.NewJwt().IssueToken(user.GetStringId(), user.Name)
		response.Json(c, gin.H{
			"token": token,
		})
	}
}
