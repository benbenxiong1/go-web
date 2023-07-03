package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-web/pkg/console"
	"os"
)

var CmdMakeModel = &cobra.Command{
	Use:   "model",
	Short: "Crate model file, example: make model user",
	Run:   runMakeModel,
	Args:  cobra.ExactArgs(1),
}

func runMakeModel(cmd *cobra.Command, args []string) {
	// 格式化模型名称， 返回一个 Model 对象
	model := makeModelFromString(args[0])

	// 确保模型的目录存在，例如 `app\model\user`
	dir := fmt.Sprintf("app/model/%s/", model.PackageName)

	// os.MkdirAll 会确保父目录和子目录都会创建，第二个参数是目录权限，使用0777
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		console.Exit(err.Error())
	}

	// 替换变量
	createFileFromStub(dir+model.PackageName+"_model.go", "model/model", model)
	createFileFromStub(dir+model.PackageName+"_util.go", "model/model_util", model)
	createFileFromStub(dir+model.PackageName+"_hooks.go", "model/model_hooks", model)
}
