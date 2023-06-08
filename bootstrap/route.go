// Package bootstrap  user : benbenxiong  time : 2023-06-2023/6/6 21:21:47
package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go-web/app/http/middlewares"
	"go-web/routes"
	"net/http"
	"strings"
)

func SetupRoute(route *gin.Engine) {
	// 注册全部中间件
	registerGlobalMiddleWare(route)
	// 注册api路由
	routes.RegisterAPIRoutes(route)
	// 配置404路由
	setup404Handler(route)
}

func registerGlobalMiddleWare(route *gin.Engine) {
	route.Use(middlewares.Logger(), middlewares.Recovery())
}

func setup404Handler(route *gin.Engine) {
	// 处理404请求
	route.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
