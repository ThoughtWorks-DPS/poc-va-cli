package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

func init() {
	rootCmd.AddCommand(teamCmd)
}

var teamCmd = &cobra.Command{
	Use: "team",
	Short: "Hit an endpoint",
	Long: "Long description to hit an endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		httpBinEndPoint := SomeEndPoint{URL: "http://localhost:3000/get"}
		HitEndPoint(httpBinEndPoint.URL)
	},
}

type SomeEndPoint struct {
	URL string
}

func HitEndPoint(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Could not connect to httpbin. Reason: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Successfully hit httpbin.org")
		return "Successfully hit httpbin.org"
	} else {
		fmt.Println("httpbin returned an error: ", resp.StatusCode)
		return fmt.Sprintf("httpbin returned an http status error: %d", resp.StatusCode)
	}
	return ""
}