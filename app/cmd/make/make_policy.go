package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-web/pkg/console"
	"os"
)

var CmdMakePolicy = &cobra.Command{
	Use:   "policy",
	Short: "",
	Run:   runMakePolicy,
	Args:  cobra.ExactArgs(1),
}

func runMakePolicy(cmd *cobra.Command, args []string) {
	// 实例化模型名称
	model := makeModelFromString(args[0])

	// os.MkdirAll 会确保父目录和子目录都被创建 第二个参数是目录权限 使用0777
	err := os.MkdirAll("app/policies", os.ModePerm)
	if err != nil {
		console.Error(err.Error())
	}

	// 拼接目标文件路径
	filePath := fmt.Sprintf("app/policies/%s_policy.go", model.PackageName)

	// 基于模板创建文件
	createFileFromStub(filePath, "policy", model)
}