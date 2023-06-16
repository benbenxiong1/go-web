package cmd

import (
	"github.com/spf13/cobra"
	"go-web/pkg/console"
	"go-web/pkg/helpers"
)

var CmdKey = &cobra.Command{
	Use:   "key",
	Short: "Generate App Key, will print the generated Key",
	Run:   runKeyGenerate,
	Args:  cobra.NoArgs,
}

func runKeyGenerate(cmd *cobra.Command, args []string) {
	console.Success("<-------- App Key: -------->")
	console.Success(helpers.RandomString(32))
	console.Success("<-------- App Key: -------->")
	console.Warning("please go to .env file to change the APP_KEY option")
}
