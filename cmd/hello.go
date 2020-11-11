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
		doHello(service)
	},
}

type ApiUrl struct {
	URL string
}
func doHello(s Service) string {
	apiUrl := ApiUrl{URL: os.Getenv("URL") + "/hello"}.URL
	status := s.getHello(apiUrl)

	fmt.Println(status)
	return status
}

type Service interface {
	getHello(url string) string
}

type HelloService struct {
}

func (s HelloService) getHello(url string) string {
	res, _ := http.Get(url)
	if(res.StatusCode == 200){
		defer res.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		bodyString := string(bodyBytes)
		return bodyString
	}
	return "Could not reach API"
}