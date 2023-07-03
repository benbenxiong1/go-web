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

var CmdMigrateDown = &cobra.Command{
	Use:     "down",
	Aliases: []string{"rollback"}, // 设置别名 migrate down == migrate rollback
	Short:   "",
	Run:     runDown,
}

var CmdMigrateReset = &cobra.Command{
	Use:   "reset",
	Short: "Rollback all database migrations",
	Run:   runReset,
}

var CmdMigrateRefresh = &cobra.Command{
	Use:   "refresh",
	Short: "Reset and re-run all migrations",
	Run:   runRefresh,
}

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,      // 执行迁移
		CmdMigrateDown,    // 回滚上一次迁移
		CmdMigrateReset,   // 回滚所有迁移
		CmdMigrateRefresh, // 回滚所有迁移 并执行所有迁移
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

func runDown(cmd *cobra.Command, args []string) {
	migrator().Rollback()
}

func runReset(cmd *cobra.Command, args []string) {
	migrator().Reset()
}

func runRefresh(cmd *cobra.Command, args []string) {
	migrator().Refresh()
}
