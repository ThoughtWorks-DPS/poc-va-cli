package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

func init() {
	rootCmd.AddCommand(helloCmd)
}

var helloCmd = &cobra.Command{
	Use: "hello",
	Short: "Call hello endpoints in API",
	Long: "Call to the API endpoint /hello",
	Run: func(cmd *cobra.Command, args []string) {
		service := HelloService{url: "https://httpbin.org/get"}
		hello(service)
	},
}

func hello(s Service) string {
	status := s.hello()

	fmt.Println(status)
	return status
}

type Service interface {
	hello() string
}

type HelloService struct {
	url string
}

func (s HelloService) hello() string {
	res, _ := http.Get(s.url)
	return res.Status
}
