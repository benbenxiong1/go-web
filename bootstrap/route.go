package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go-web/app/http/middlewares"
	"go-web/pkg/response"
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
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
		)
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
			response.Abort404(c)
		}


	})
}