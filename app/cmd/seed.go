package cmd

import (
	"github.com/spf13/cobra"
	"go-web/database/seeders"
	"go-web/pkg/console"
	"go-web/pkg/seed"
)

var CmdDBSeeder = &cobra.Command{
	Use:   "seed",
	Short: "",
	Run:   runDBSeeder,
	Args:  cobra.MaximumNArgs(1),
}

func runDBSeeder(cmd *cobra.Command, args []string) {
	seeders.Initialize()
	if len(args) > 0 {
		// 有传参数的情况
		name := args[0]
		seeder := seed.GetSeeder(name)
		if len(seeder.Name) > 0 {
			seed.RunSeeder(name)
		} else {
			console.Error("Seeder not found: " + name)
		}
	} else {
		// 运行所有的
		seed.RunAll()
		console.Success("Done seeding.")
	}

}
