package cmd

import (
	"github.com/spf13/cobra"
	"go-web/database/migrations"
	"go-web/pkg/migrate"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "",
}

var CmdMigrateUp = &cobra.Command{
	Use:   "up",
	Short: "",
	Run:   runUp,
}

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
	)
}

func migrator() *migrate.Migrator {
	// 注册 database/migrations 下所有的迁移文件
	migrations.Initialize()

	// 初始化 migrator
	return migrate.NewMigrator()
}

func runUp(cmd *cobra.Command, args []string) {
	migrator().Up()
}
