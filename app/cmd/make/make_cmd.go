package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-web/pkg/console"
)

var CmdMakeCmd = &cobra.Command{
	Use:   "cmd",
	Short: "Create a command, should be snake_case, exmaple: make cmd buckup_database",
	Run:   runMakeCmd,
	Args:  cobra.ExactArgs(1), // 只允许且必须传1个参数
}

func runMakeCmd(cmd *cobra.Command, args []string) {
	// 格式化模型名称，返回一个 Model 对象
	model := makeModelFromString(args[0])

	// 拼接目标文件路径
	filepath := fmt.Sprintf("app/cmd/%s.go", model.PackageName)

	// 从模板中创建文件(做好变量替换)
	createFileFromStub(filepath, "cmd", model)

	// 友好提示
	console.Success("command name:" + model.PackageName)
	console.Success("command variable name:cmd.Cmd" + model.StructName)
	console.Success("please edit main.go's app.Commands slice to register command")
}
