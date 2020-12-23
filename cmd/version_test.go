package cmd

import (
	"github.com/kami-zh/go-capturer"
	"strings"
	"testing"
)

func ExamplePrintCliVersionWithDefaults() {
	printCliVersion()
	//Output:
	//voltron version: dev, SHA: dev
}

func TestVersionCmdIntegration(t *testing.T) {
	cmd := versionCmd()
	out := capturer.CaptureStdout(func() {
		cmd.ExecuteC()
	})

	if !strings.Contains(string(out), "voltron version: dev, SHA: dev") {
		t.Fatalf("expected \"%s\" got \"%s\"", "voltron version: dev, SHA: dev", string(out))
	}
}