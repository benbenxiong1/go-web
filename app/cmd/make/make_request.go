package make

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CmdMakeRequest = &cobra.Command{
	Use:   "request",
	Short: "Create request file, example make request user",
	Run:   runMakeRequest,
	Args:  cobra.ExactArgs(1),
}

func runMakeRequest(cmd *cobra.Command, args []string) {
	// 格式化模型名称，返回一个 model 对象
	model := makeModelFromString(args[0])

	// 拼接目标路径
	filePath := fmt.Sprintf("app/requests/%s_request.go", model.PackageName)

	// 基于模板创建文件
	createFileFromStub(filePath, "request", model)
}
