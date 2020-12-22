package cmd

import (
	"github.com/kami-zh/go-capturer"
	"strings"
	"testing"
)

func ExampleTestVersionWithDefaults() {
	printVersion()
	//Output:
	//voltron version: development, SHA: gitHash
}

func TestVersionCmdIntegration(t *testing.T) {
	cmd := versionCmd()
	out := capturer.CaptureStdout(func() {
		cmd.ExecuteC()
	})

	if !strings.Contains(string(out), "voltron version: development, SHA: gitHash") {
		t.Fatalf("expected \"%s\" got \"%s\"", "voltron version: development, SHA: gitHash", string(out))
	}
}