package cmd

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"voltron/clients"
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

