package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"voltron/clients"
)

func init() {
	rootCmd.AddCommand(versionCmd())
}

func versionCmd() *cobra.Command {
  return &cobra.Command{
  	Use: "version",
  	Short: "Current version of Voltron and API",
  	Run: func(cmd *cobra.Command, args []string) {
		printVersions()
	},
  }
}

var Version = "dev"
var Commit = "dev"

func printVersions() {
	printCliVersion()
	printApiVersion()
}

func printCliVersion() {
	fmt.Printf("voltron version: %s, SHA: %s\n", Version, Commit)
}

func printApiVersion() {
	client := clients.NewApiClient()
	info := client.GetApiInfo()
	fmt.Println(info)
}