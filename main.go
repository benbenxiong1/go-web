package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/bootstrap"
	bsConfig "go-web/config"
	"go-web/pkg/config"
)

func init() {
	bsConfig.Initialize()
}

func main()  {

	var env string
	flag.StringVar(&env,"env","","加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")

	flag.Parse()
	config.InitConfig(env)

	// new 一个 Gin Engine 实例
	r := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(r)

	// 运行服务
	err := r.Run(":" + config.Get("app.port"))
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
		return
	}
}