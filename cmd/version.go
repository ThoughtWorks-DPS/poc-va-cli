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
		client := clients.NewApiClient()
		printCliVersion()
		fmt.Println(getApiVersion(client))
	},
  }
}

var Version = "dev"
var Commit = "dev"

func printCliVersion() {
	fmt.Printf("voltron version: %s, SHA: %s\n", Version, Commit)
}

func getApiVersion(client clients.ApiClient) string {
	return client.GetApiInfo()
}