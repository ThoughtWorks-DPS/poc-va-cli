package clients

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

type ApiClient interface {
	GetHello() string
}

type ApiClientImpl struct {
	URL string
}

func NewApiClient() ApiClientImpl {
	client := ApiClientImpl{}
	client.URL = viper.GetString("API_SERVICE_BASE_URL")
	return client
}

func (client ApiClientImpl) GetHello() string {
	helloEndPoint := client.URL + "/teams/hello"
	res, _ := http.Get(helloEndPoint)

	if res.StatusCode == 200 {
		defer res.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		bodyString := string(bodyBytes)
		return bodyString
	}

	return "Could not reach API"
}