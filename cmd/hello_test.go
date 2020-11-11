package cmd

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockedService struct{
	mock.Mock
}

func (m MockedService) getHello(url string) string {
	url = "www.example.com/hello"
	args := m.Called(url)
	return args.String(0)
}

func TestDoHello(t *testing.T) {
	mockedService := MockedService{}
	mockedService.On("getHello", "www.example.com/hello").Return("Hello from the API!", nil)

	actual := doHello(mockedService)

	assert.Equal(t, actual, "Hello from the API!")
}

func TestGetHelloSuccessfulCall(t *testing.T) {
	service := HelloService{}
	 var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	 	w.WriteHeader(http.StatusOK)
	 	w.Write([]byte("Hello from the API!"))
	}))

	assert.Equal(t, "Hello from the API!", service.getHello(apiStub.URL))
}

func TestGetHelloFailedCall(t *testing.T) {
	service := HelloService{}
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))

	assert.Equal(t, "Could not reach API", service.getHello(apiStub.URL))
}