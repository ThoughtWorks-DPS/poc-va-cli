package clients

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"

	"testing"
)

func TestGetHelloSuccessfulCall(t *testing.T) {
	client := NewApiClient()
	var apiStub = httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(http.StatusOK)
		response.Write([]byte("Hello from the API!"))

		assert.Equal(t, "/teams/hello", request.URL.Path)
	}))

	client.URL = apiStub.URL

	assert.Equal(t, "Hello from the API!", client.GetHello())
}

func TestGetHelloFailedCall(t *testing.T) {
	client := NewApiClient()
	var apiStub = httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(http.StatusInternalServerError)
	}))

	client.URL = apiStub.URL

	assert.Equal(t, "Error: Service Returned 500 Internal Server Error", client.GetHello())
}

func TestSetDefaultApiUrl(t *testing.T) {
	loadViperConfig()
	client := NewApiClient()
	assert.Equal(t, "http://api.devportal.name", client.URL )
}

func TestOverrideDefaultApiUrl(t *testing.T) {
	client := NewApiClient()
	os.Setenv("API_SERVICE_BASE_URL", "http://localhost:5000")

	assert.Equal(t, "http://localhost:5000", client.URL)
}

func loadViperConfig() {
	viper.AddConfigPath("../")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}