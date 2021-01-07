package cmd

import (
	"fmt"
	"strings"
	"testing"
	"voltron/mocks"

	"github.com/golang/mock/gomock"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func ExampleTestPrintCliDefaultVersion() {
	printCliVersion()
	//Output:
	//voltron version: dev, SHA: dev
}

func TestShortenCommitHashGreaterThan7(t *testing.T) {
	hash := "12345678"
	assert.Equal(t, "1234567", shortenCommitHash(hash))
}

func TestShortenCommitHashLessThan7(t *testing.T) {
	hash := "123456"
	assert.Equal(t, "123456", shortenCommitHash(hash))
}

func TestGetApiVersion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedApiClient := mocks.NewMockApiClient(ctrl)
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
