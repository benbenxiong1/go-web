package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main()  {

	r := gin.New()

	r.Use(gin.Logger(),gin.Recovery())


	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"Hello Word",
		})
	})

	r.NoRoute(func(c *gin.Context) {
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

	r.Run(":8000")

	fmt.Println("Hello 世界!")
}