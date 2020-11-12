package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
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
		client := HelloClient{}
		doHello(client)
	},
}

type ApiUrl struct {
	URL string
}

type RestClient interface {
	getHello(url string) string
}

type HelloClient struct {
}

func doHello(client RestClient) string {
	apiUrl := ApiUrl{URL: viper.GetString("API_SERVICE_BASE_URL") + "/hello"}
	status := client.getHello(apiUrl.URL)

	fmt.Println(status)
	return status
}

func (client HelloClient) getHello(url string) string {
	res, _ := http.Get(url)
	if res.StatusCode == 200 {
		defer res.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		bodyString := string(bodyBytes)
		return bodyString
	}
	return "Could not reach API"
}