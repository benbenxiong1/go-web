package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-web/pkg/app"
	"go-web/pkg/console"
)

var CmdMakeMigration = &cobra.Command{
	Use:   "migration",
	Short: "",
	Run:   runMakeMigration,
	Args:  cobra.ExactArgs(1),
}

func runMakeMigration(cmd *cobra.Command, args []string) {
	// 日期格式化
	timeStr := app.TImeNowInTimezone().Format("2006_01_02_150405")

	model := makeModelFromString(args[0])

	fileName := timeStr + "_" + model.PackageName
	filePath := fmt.Sprintf("database/migrations/%s.go", fileName)
	createFileFromStub(filePath, "migration", model, map[string]string{
		"{{FileName}}": fileName})
	console.Success("Migration file created，after modify it, use `migrate up` to migrate database.")

}
