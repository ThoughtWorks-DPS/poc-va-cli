package cmd

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestHelloWorld(t *testing.T){
	service := HelloService{"https://httpbin.org/get"}
	s := hello(service)

	assert.Equal(t, "200 OK", s, "Status code should equal 200 OK" )
}


type MockedService struct{
	mock.Mock
}

func (m MockedService) hello() string {
	args := m.Called()
	return args.String(0)

}

func TestCallWasCalled(t *testing.T) {
	mockedService := MockedService{}
	mockedService.On("hello", ).Return("200 OK", nil)

	hello(mockedService)

	mockedService.AssertExpectations(t)
}