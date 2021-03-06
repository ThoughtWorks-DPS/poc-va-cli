package clients

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

type ApiClient interface {
	GetHello() string
	GetApiInfo() string
}

type ApiClientImpl struct {
	URL string
}

type ApiInfo struct {
	Application struct {
		SemVersion string `json:"SemVersion"`
		GitHash    string `json:"GitHash"`
	} `json:"application"`
}

func NewApiClient() ApiClientImpl {
	client := ApiClientImpl{}
	client.URL = viper.GetString("api_service_base_url")
	return client
}

func (client ApiClientImpl) GetHello() string {
	helloEndPoint := client.URL + "/teams/hello"
	res, err := http.Get(helloEndPoint)

	if err != nil {
		return "Invalid url."
	}

	if res.StatusCode == 200 {
		defer res.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		bodyString := string(bodyBytes)
		return bodyString
	}

	return "Error: API Service Returned " + res.Status
}

func (client ApiClientImpl) GetApiInfo() string {
	infoEndPoint := client.URL + "/teams/info"
	res, err := http.Get(infoEndPoint)
	info := ApiInfo{}
	if err != nil {
		return "Invalid url."
	}

	if res.StatusCode == 200 {
		defer res.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		err := json.Unmarshal(bodyBytes, &info)

		if err != nil {
			log.Fatal(err)
		}

		return "API version: " + info.Application.SemVersion + ", SHA: " + info.Application.GitHash + ", Base URL: " + client.URL
	}

	return "Error: API Service Returned " + res.Status
}
