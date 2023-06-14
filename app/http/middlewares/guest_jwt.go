package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-web/pkg/jwt"
	"go-web/pkg/response"
)

func GuestJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) > 0 {

			_, err := jwt.NewJwt().ParserToken(c)
			if err == nil {
				response.Unauthorized(c, "请使用游客身份访问")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
