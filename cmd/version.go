package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd())
}

func versionCmd() *cobra.Command {
  return &cobra.Command{
  	Use: "version",
  	Short: "Current version of Voltron and API",
  	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
  }
}

var Version = "development"
var Commit = "gitHash"

func printVersion() {
	fmt.Printf("voltron version: %s, SHA: %s\n", Version, Commit)
}