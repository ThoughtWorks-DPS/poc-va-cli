package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"io/ioutil"
)

func init() {
	rootCmd.AddCommand(helloCmd)
}

var helloCmd = &cobra.Command{
	Use: "hello",
	Short: "Call hello endpoints in API",
	Long: "Call to the API endpoint /hello",
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Getenv("URL")) == 0 {
			os.Setenv("URL", "https://devportal.name/teams")
		}
		service := HelloService{}
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
}

func (s HelloService) hello() string {
	res, _ := http.Get(os.Getenv("URL") + "/hello")
	defer res.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	bodyString := string(bodyBytes)
	return bodyString
}