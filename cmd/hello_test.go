package cmd

import (
	"github.com/ThoughtWorks-DPS/poc-va-cli/clients"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDoHello(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedApiClient := clients.NewMockApiClient(ctrl)
	mockedApiClient.EXPECT().
		GetHello().
		Return("Hello from the API!")

	assert.Equal(t, "Hello from the API!", doHello(mockedApiClient))
}

func TestGetHelloSuccessfulCall(t *testing.T) {
	client := clients.NewApiClient()
	var apiStub = httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(http.StatusOK)
	 	response.Write([]byte("Hello from the API!"))

		assert.Equal(t, "/teams/hello", request.URL.Path)
	}))

	 client.URL = apiStub.URL

	 assert.Equal(t, "Hello from the API!", client.GetHello())
}

func TestGetHelloFailedCall(t *testing.T) {
	client := clients.NewApiClient()
	var apiStub = httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(http.StatusInternalServerError)
	}))

	client.URL = apiStub.URL

	assert.Equal(t, "Could not reach API", client.GetHello())
}
