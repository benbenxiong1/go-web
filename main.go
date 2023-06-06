// Package go_web  user : benbenxiong  time : 2023-06-2023/6/6 22:22:07
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/bootstrap"
)

func main() {
	// new一个 gin.Engine 实例
	r := gin.New()

	// 初始化路由
	bootstrap.SetupRoute(r)

	// 运行服务
	err := r.Run(":8000")
	if err != nil {
		fmt.Println("route run err", err.Error())
	}
}
