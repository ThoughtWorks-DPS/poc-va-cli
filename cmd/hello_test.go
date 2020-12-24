package cmd

import (
	"github.com/golang/mock/gomock"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"voltron/mocks"
)

func TestDoHello(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedApiClient := mocks.NewMockApiClient(ctrl)
	mockedApiClient.EXPECT().
		GetHello().
		Return("Hello from the AP!")

	assert.Equal(t, "Hello from the API!", doHello(mockedApiClient))
}

func Test_ExecuteHelloCommand(t *testing.T) {
	cmd := helloCmd()
	out := capturer.CaptureStdout(func() {
		cmd.ExecuteC()
	})

	if !strings.Contains(string(out), "Hello from the API!") {
		t.Fatalf("expected \"%s\" got \"%s\"", "Hello from the API!", string(out))
	}
}
