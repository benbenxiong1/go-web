package make

import (
	"embed"
	"fmt"
	"github.com/spf13/cobra"
	"go-web/pkg/console"
	"go-web/pkg/file"
	"go-web/pkg/str"
	"strings"
)

// Model 参数解释
// 单个词，用户命令传参，以 User 模型为例：
//     - user
//  - User
//  - users
//  - Users
// 整理好的数据：
// {
//     "TableName": "users",
//     "StructName": "User",
//     "StructNamePlural": "Users"
//     "VariableName": "user",
//     "VariableNamePlural": "users",
//     "PackageName": "user"
// }
// -
// 两个词或者以上，用户命令传参，以 TopicComment 模型为例：
//     - topic_comment
//  - topic_comments
//  - TopicComment
//  - TopicComments
// 整理好的数据：
// {
//     "TableName": "topic_comments",
//     "StructName": "TopicComment",
//     "StructNamePlural": "TopicComments"
//     "VariableName": "topicComment",
//     "VariableNamePlural": "topicComments",
//     "PackageName": "topic_comment"
// }

type Model struct {
	TableName          string
	StructName         string
	StructNamePlural   string
	VariableName       string
	VariableNamePlural string
	PackageName        string
}

// stubsFS 方便我们后面打包这些 .stub 为后缀名的文件
//go:embed stubs
var stubsFs embed.FS

var CmdMake = &cobra.Command{
	Use:   "make",
	Short: "Generate file and code",
}

func init() {
	// 注册子命令
	CmdMake.AddCommand(
		CmdMakeCmd,           // 生成自定义命令
		CmdMakeModel,         // 生成模型文件
		CmdMakeApiController, // 生成控制器
	)
}

// makeModelFromString 格式化用户输入的内容
func makeModelFromString(name string) Model {
	return Model{
		StructName:         str.Singular(name),
		StructNamePlural:   str.Plural(name),
		TableName:          str.Snake(name),
		VariableName:       str.LowerCamel(name),
		PackageName:        str.Snake(name),
		VariableNamePlural: str.LowerCamel(name),
	}
}

// createFileFromStub 读取 stub 文件并进行变量替换
func createFileFromStub(filePath string, stubName string, model Model, variables ...interface{}) {
	// 实现最后一个参数可选
	replaces := make(map[string]string)
	if len(variables) > 0 {
		replaces = variables[0].(map[string]string)
	}

	// 判断文件是否存在
	if file.Exists(filePath) {
		console.Exit(filePath + " already exists!")
	}

	// 读取 stub 模板文件
	modelData, err := stubsFs.ReadFile("stubs/" + stubName + ".stub")
	if err != nil {
		console.Exit(err.Error())
	}
	modelStub := string(modelData)

	// 添加默认替换的变量
	replaces["{{VariableName}}"] = model.VariableName
	replaces["{{VariableNamePlural}}"] = model.VariableNamePlural
	replaces["{{StructName}}"] = model.StructName
	replaces["{{StructNamePlural}}"] = model.StructNamePlural
	replaces["{{PackageName}}"] = model.PackageName
	replaces["{{TableName}}"] = model.TableName

	// 对模板内容做变量替换
	for search, replace := range replaces {
		modelStub = strings.ReplaceAll(modelStub, search, replace)
	}

	// 存储到目标文件中
	err = file.Put([]byte(modelStub), filePath)
	if err != nil {
		console.Exit(err.Error())
	}

	// 提示成功
	console.Success(fmt.Sprintf("[%s] creeated.", filePath))
}