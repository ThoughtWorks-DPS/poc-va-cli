package cmd

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

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