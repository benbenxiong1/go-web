package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/app/model/user"
	"go-web/pkg/config"
	"go-web/pkg/jwt"
	"go-web/pkg/response"
)

// AuthJwt jwt 登录中间件
// 注意中间件里，当在 c.Next() 之前 return 掉，就会中断所有的后续请求。
func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		cliams, err := jwt.NewJwt().ParserToken(c)
		// JWT 解析失败，有错误发生
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}

		// 解析成功获取用户信息
		userModel := user.Get(cliams.UserId)
		if userModel.ID == 0 {
			response.Unauthorized(c, "找不到对应的用户，用户可能已删除")
			return
		}

		// 将用户信息存入 gin.context里 后续auth包将从这里拿当前用户数据
		c.Set("current_user_id", userModel.ID)
		c.Set("current_user_name", userModel.Name)
		c.Set("current_user", userModel)

		c.Next()
	}
}
