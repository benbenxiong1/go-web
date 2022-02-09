package routes

import (
	"github.com/gin-gonic/gin"
	"go-web/app/http/controllers/api/v1/auth"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(route *gin.Engine) {
	v1 := route.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			signup := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", signup.IsPhoneExist)
			// 判断email是否已注册
			authGroup.POST("/signup/email/exist", signup.IsEmailExist)
			//手机号注册用户
			authGroup.POST("/signup/using-phone", signup.SignupUsingPhone)
			//邮箱注册用户
			authGroup.POST("/signup/using-email", signup.SignupUsingEmail)

			// 发送验证码
			verifyCode := new(auth.VerifyCodeController)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-codes/captcha", verifyCode.ShowCaptcha)
			//短信验证码
			authGroup.POST("/verify-codes/phone", verifyCode.SendUsingPhone)
			//发送邮件
			authGroup.POST("/verify-codes/email", verifyCode.SendUsingEmail)

		}

	}
}
