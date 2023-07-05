// Package routes  user : benbenxiong  time : 2023-06-2023/6/6 22:22:15
package routes

import (
	"github.com/gin-gonic/gin"
	v12 "go-web/app/http/controllers/api/v1"
	"go-web/app/http/controllers/api/v1/auth"
	"go-web/app/http/middlewares"
)

func RegisterAPIRoutes(route *gin.Engine) {

	v1 := route.Group("v1")
	v1.Use(middlewares.LimitIp("1000-H"))
	{
		authGroup := v1.Group("/auth")
		authGroup.Use(middlewares.LimitIp("200-H"))
		{
			// 注册
			suc := new(auth.SignupController)
			// 验证手机号是否存在
			authGroup.POST("/signup/phone/exist", middlewares.GuestJwt(), suc.IsPhoneExist)
			// 验证邮箱是否存在
			authGroup.POST("/signup/email/exist", middlewares.GuestJwt(), suc.IsEmailExist)
			// 手机号注册
			authGroup.POST("/signup/using-phone", middlewares.GuestJwt(), suc.SignupUsingPhone)
			// 邮箱注册
			authGroup.POST("/signup/using-email", middlewares.GuestJwt(), suc.SignupUsingEmail)

			// 验证
			vcc := new(auth.VerifyCodeController)
			// 验证码
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)
			// 验证手机号+验证码
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
			// 验证email+验证码
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)

		}

		// 登录
		login := new(auth.LoginController)
		// 手机号 + 验证码登录
		authGroup.POST("/login/using-phone", middlewares.GuestJwt(), login.LoginByPhone)
		// 支持手机号/用户名/邮箱登录
		authGroup.POST("/login/using-password", middlewares.GuestJwt(), login.Login)
		// 刷新token
		authGroup.POST("/login/refresh-token", middlewares.AuthJwt(), login.RefreshToken)

		// 修改密码
		pwd := new(auth.PasswordController)
		// 手机号+验证码修改密码
		authGroup.POST("/password-reset/using-phone", middlewares.LimitPerRoute("5-H"), pwd.ResetByPhone)
		// 邮箱+验证码修改密码
		authGroup.POST("/password-reset/using-email", middlewares.LimitPerRoute("5-H"), pwd.ResetByEmail)

		user := new(v12.UsersController)
		// 获取当前用户
		v1.GET("/user", middlewares.AuthJwt(), user.CurrenUser)
		userGroup := v1.Group("/users")
		{
			// 获取所有用户
			userGroup.GET("", user.Index)
		}

		// 分类
		cgc := new(v12.CategoriesController)
		cgcGroup := v1.Group("/categories")
		{
			// 列表
			cgcGroup.GET("", cgc.Index)
			// 新增
			cgcGroup.POST("", cgc.Store)
			// 修改
			cgcGroup.PUT("/:id", cgc.Update)
			// 删除
			cgcGroup.DELETE("/:id", cgc.Delete)
		}
	}
}
