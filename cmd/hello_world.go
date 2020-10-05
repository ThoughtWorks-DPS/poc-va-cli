package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(printCmd)
}

var printCmd = &cobra.Command{
	Use: "print",
	Short: "print hello world",
	Long: "print long description of print command.",
	Run: func(cmd *cobra.Command, args []string) {
		helloWorld()
	},
}

func helloWorld() string {
	fmt.Println("Hello world")
	return "Hello world"
}