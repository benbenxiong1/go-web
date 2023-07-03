package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-web/pkg/console"
	"os"
	"strings"
)

var CmdMakeApiController = &cobra.Command{
	Use:   "apiController",
	Short: "",
	Run:   runMakeApiController,
	Args:  cobra.ExactArgs(1),
}

func runMakeApiController(cmd *cobra.Command, args []string) {
	// 处理参数，要求附带 API 版本 （v1 或者 v2）
	array := strings.Split(args[0], "/") // 按 / 拆分
	if len(array) != 2 {
		console.Exit("api controller name format:v1/user")
	}

	// apiVersion 用来拼接目标路径
	// name 用来生成 cmd.model 实例
	apiVersion, name := array[0], array[1]
	model := makeModelFromString(name)

	// 确保目录存在
	dir := fmt.Sprintf("app/http/controllers/api/%s", apiVersion)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		console.Exit(err.Error())
	}

	// 组建目标目录
	filePath := fmt.Sprintf(dir+"/%s_conntroller.go", model.TableName)

	// 基于模板创建文件 （做好变量替换）
	createFileFromStub(filePath, "apicontroller", model, map[string]string{
		"{{version}}": apiVersion})
}
