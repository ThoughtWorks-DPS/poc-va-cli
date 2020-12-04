package cmd

import (
	"testing"
	"voltron/clients"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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

func TestCodeClimate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedApiClient := clients.NewMockApiClient(ctrl)
	mockedApiClient.EXPECT().
		GetHello().
		Return("Hello from the API!")

	assert.Equal(t, "Failing test to validate code climate story", doHello(mockedApiClient))
}
