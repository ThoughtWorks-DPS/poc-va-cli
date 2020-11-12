package cmd

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)


func TestDoHelloDefaultApiUrl(t *testing.T) {
	setUpViper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockedHelloClient := NewMockRestClient(ctrl)

	mockedHelloClient.EXPECT().
		getHello("http://api.devportal.name/hello").
		Return("Hello from the API!")

	status := doHello(mockedHelloClient)

	assert.Equal(t, "Hello from the API!", status)
}

func TestDoHelloOverrideApiUrl(t *testing.T) {
	setUpViper()
	os.Setenv("API_SERVICE_BASE_URL", "http://localhost:5000")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedHelloClient := NewMockRestClient(ctrl)
	mockedHelloClient.EXPECT().getHello("http://localhost:5000/hello")

	doHello(mockedHelloClient)
}

func TestGetHelloSuccessfulCall(t *testing.T) {
	client := HelloClient{}
	 var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	 	w.WriteHeader(http.StatusOK)
	 	w.Write([]byte("Hello from the API!"))
	}))

	assert.Equal(t, "Hello from the API!", client.getHello(apiStub.URL))
}

func TestGetHelloFailedCall(t *testing.T) {
	client := HelloClient{}
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))

	assert.Equal(t, "Could not reach API", client.getHello(apiStub.URL))
}

func setUpViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath("..")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}