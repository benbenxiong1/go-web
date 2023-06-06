// Package routes  user : benbenxiong  time : 2023-06-2023/6/6 22:22:15
package routes

import (
	"github.com/gin-gonic/gin"
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
	}
}
