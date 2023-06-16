// Package go_web  user : benbenxiong  time : 2023-06-2023/6/6 22:22:07
package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-web/app/cmd"
	"go-web/bootstrap"
	baseConfig "go-web/config"
	"go-web/pkg/config"
	"go-web/pkg/console"
	"os"
)

func init() {
	// 加载config下的配置信息
	baseConfig.Initialize()
}

func main() {

	var rootCmd = &cobra.Command{
		Use:   "Gohub",
		Short: "A simple forum project",
		Long:  `Default will run "serve" command ,you can use "-h" flag to see all subcommands`,
		PersistentPostRun: func(Command *cobra.Command, args []string) {
			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)

			// 初始化logger
			bootstrap.SetupLogger()

			// 初始化数据库
			bootstrap.SetupDB()

			// 初始化redis
			bootstrap.SetupRedis()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
	)

	// 配置默认运行 web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数 --env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}

	//var env string
	//flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	//flag.Parse()
	//config.InitConfig(env)
	//
	//// 初始化日志
	//bootstrap.SetupLogger()
	//
	//// 设置 gin 的运行模式，支持 debug, release, test
	//// release 会屏蔽调试信息，官方建议生产环境中使用
	//// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	//// 故此设置为 release，有特殊情况手动改为 debug 即可
	//gin.SetMode(gin.ReleaseMode)
	//
	//// new一个 gin.Engine 实例
	//r := gin.New()
	//
	//// 初始化Redis
	//bootstrap.SetupRedis()
	//
	//// 初始化DB
	//bootstrap.SetupDB()
	//
	//// 初始化路由
	//bootstrap.SetupRoute(r)
	//
	//// 运行服务
	//err := r.Run(":" + config.Get("app.port"))
	//if err != nil {
	//	fmt.Println("route run err", err.Error())
	//}
}
