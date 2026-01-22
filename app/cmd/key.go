package cmd

import (
	"github.com/Zero0719/go-api/helpers"
	"github.com/Zero0719/go-api/pkg/console"
	"github.com/spf13/cobra"
)

var CmdKey = &cobra.Command{
	Use: "key",
	Short: "Generate app key",
	Run: runKeyGenerate,
	Args: cobra.NoArgs,
}

func runKeyGenerate(cmd *cobra.Command, args []string) {
	console.Success("---")
    console.Success("App Key:")
    console.Success(helpers.RandomString(32))
    console.Success("---")
    console.Warning("please go to .env file to change the APP_KEY option")
}