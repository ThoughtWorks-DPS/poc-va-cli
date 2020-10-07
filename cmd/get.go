package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use: "get",
	Short: "call a get command",
	Long: "print long description of print command.",
	Run: func(cmd *cobra.Command, args []string) {
		get()
	},
}

func get() string {
	res, _ := http.Get("https://httpbin.org/get")

	fmt.Println(res.Status)
	return res.Status

}