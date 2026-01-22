package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdPlay = &cobra.Command{
	Use: "play",
	Short: "Play with the API",
	Run: runPlay,
	Args: cobra.NoArgs,
}

func runPlay(cmd *cobra.Command, args []string) {
	fmt.Println("Hello, World!")
}