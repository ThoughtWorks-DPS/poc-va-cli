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
	mockedHelloClient := clients.NewMockRestClient(ctrl)

	mockedHelloClient.EXPECT().
		GetHello().
		Return("Hello from the API!")

	status := doHello(mockedHelloClient)

	assert.Equal(t, "Hello from the API!", status)
}

func TestGetHelloSuccessfulCall(t *testing.T) {
	client := clients.NewApiClient()

	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	 	w.Write([]byte("Hello from the API!"))
	}))

	 client.URL = apiStub.URL

	 assert.Equal(t, "Hello from the API!", client.GetHello())
}

func TestGetHelloFailedCall(t *testing.T) {
	client := clients.NewApiClient()
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))

	client.URL = apiStub.URL

	assert.Equal(t, "Could not reach API", client.GetHello())
}
