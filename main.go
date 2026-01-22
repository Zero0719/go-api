package main

import (
	"fmt"
	"os"

	"github.com/Zero0719/go-api/app/cmd"
	"github.com/Zero0719/go-api/bootstrap"
	btsConfig "github.com/Zero0719/go-api/config"
	"github.com/Zero0719/go-api/pkg/config"
	"github.com/Zero0719/go-api/pkg/console"
	"github.com/spf13/cobra"
)

func init() {
	btsConfig.Initialize()
}

func main() {
	var rootCmd = &cobra.Command{
		Use: "GoApi",
		Short: "GoApi 是一个 Go 语言的 API 框架",
		Long: `GoApi 是一个 Go 语言的 API 框架，它可以帮助你快速构建 API 服务。`,
		PersistentPreRun: func(command *cobra.Command, args []string) {
			config.InitConfig(cmd.Env)

			bootstrap.SetupLogger()
			bootstrap.SetupDB()
			bootstrap.SetupRedis()
		},
	}

	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
	)

	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	cmd.RegisterGlobalFlags(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}