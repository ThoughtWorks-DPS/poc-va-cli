package cmd

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"voltron/clients"
)

func ExampleTestPrintCliVersion() {
	printCliVersion()
	//Output:
	//voltron version: dev, SHA: dev
}

func TestGetApiVersion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedApiClient := clients.NewMockApiClient(ctrl)
	mockedApiClient.EXPECT().
		GetApiInfo().
		Return("API version: x.x.x, SHA: a1b2c34d,")

	assert.Equal(t, "API version: x.x.x, SHA: a1b2c34d,", getApiVersion(mockedApiClient))
}

func TestVersionCmdIntegration(t *testing.T) {
	cmd := versionCmd()
	out := capturer.CaptureStdout(func() {
		cmd.ExecuteC()
	})

	fmt.Println(string(out))
	if !strings.Contains(string(out), "voltron version: dev, SHA: dev") {
		t.Fatalf("expected \"%s\" got \"%s\"", "voltron version: dev, SHA: dev", string(out))
	}

	if !strings.Contains(string(out), "API version: ") {
		t.Fatalf("expected \"%s\" got \"%s\"", "API version: x.x.x", string(out))
	}

	if !strings.Contains(string(out), "SHA: ") {
		t.Fatalf("expected \"%s\" got \"%s\"", "SHA: a1b2c3d4", string(out))
	}
}