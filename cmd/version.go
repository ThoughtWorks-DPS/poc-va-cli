package cmd

import (
	"fmt"
	"voltron/clients"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd())
}

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
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

func shortenCommitHash(hash string) string {
	if len(hash) > 7 {
		return hash[0:7]
	} else {
		return hash
	}
}

func printCliVersion() {
	fmt.Printf("voltron version: %s, SHA: %s\n", Version, shortenCommitHash(Commit))
}

func getApiVersion(client clients.ApiClient) string {
	return client.GetApiInfo()
}
