// Package routes  user : benbenxiong  time : 2023-06-2023/6/6 22:22:15
package routes

import (
	"github.com/gin-gonic/gin"
	"go-web/app/http/controllers/api/v1/auth"
	"net/http"
)

func RegisterAPIRoutes(route *gin.Engine) {

	v1 := route.Group("v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"hello": "word",
			})
		})
		authGroup := v1.Group("/auth")
		{
			// 注册
			suc := new(auth.SignupController)
			// 验证手机号是否存在
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			// 验证邮箱是否存在
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			// 手机号注册
			authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)
			// 邮箱注册
			authGroup.POST("/signup/using-email", suc.SignupUsingEmail)

			// 验证
			vcc := new(auth.VerifyCodeController)
			// 验证码
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
			// 验证手机号+验证码
			authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)
			// 验证email+验证码
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)

			// 登录
			login := new(auth.LoginController)
			//手机号 + 验证码登录
			authGroup.POST("/login/using-phone", login.LoginByPhone)
			// 支持手机号/用户名/邮箱登录
			authGroup.POST("/login/using-password", login.Login)
		}
	}
}
