package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go-web/routes"
	"net/http"
	"strings"
)

// SetupRoute 初始化路由绑定
func SetupRoute(router *gin.Engine)  {
	
	registerGlobalMiddleWare(router)

	routes.RegisterAPIRoutes(router)

	setup404Handler(router)
}

// registerGlobalMiddleWare 注册全局中间件
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(gin.Logger(),gin.Recovery())
}

// setup404Handler 处理404请求
func setup404Handler(router *gin.Engine) {

	router.NoRoute(func(c *gin.Context) {
		//获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")

		//判断是页面请求 返回页面404错误
		if strings.Contains(acceptString,"text/html"){
			c.String(http.StatusNotFound,"页面返回404")
		}else {
			c.JSON(http.StatusNotFound,gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}


	})
}