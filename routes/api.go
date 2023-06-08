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
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
		}
	}
}
