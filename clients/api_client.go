package clients

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

type RestClient interface {
	GetHello() string
}

type ApiClient struct {
	URL string
}

func NewApiClient() ApiClient {
	client := ApiClient{}
	client.URL = viper.GetString("API_SERVICE_BASE_URL")
	return client
}

func (client ApiClient) GetHello() string {
	helloEndPoint := client.URL + "/hello"
	res, _ := http.Get(helloEndPoint)
	if res.StatusCode == 200 {
		defer res.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		bodyString := string(bodyBytes)
		return bodyString
	}
	return "Could not reach API"
}