package cmd

import (
	"fmt"
	"github.com/ThoughtWorks-DPS/poc-va-cli/clients"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(helloCmd)
}

var helloCmd = &cobra.Command{
	Use: "hello",
	Short: "Call hello endpoints in API",
	Long: "Call to the API endpoint /hello",
	Run: func(cmd *cobra.Command, args []string) {
		client := clients.NewApiClient()
		doHello(client)
	},
}

func doHello(client clients.ApiClient) string {
	status := client.GetHello()
	fmt.Println(status)
	return status
}
