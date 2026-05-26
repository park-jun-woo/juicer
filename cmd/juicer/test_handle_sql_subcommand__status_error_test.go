//ff:func feature=sql type=test control=sequence
//ff:what TestHandleSQLSubcommand_StatusError 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

// TestHandleSQLSubcommand_StatusError tests that the status error branch calls os.Exit.
func TestHandleSQLSubcommand_StatusError(t *testing.T) {
	if os.Getenv("TEST_SUBPROCESS_STATUS") == "1" {
		_, cleanup := helperSetupBrokenSession(t)
		defer cleanup()
		handleSQLSubcommand([]string{"status"})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestHandleSQLSubcommand_StatusError$")
	cmd.Env = append(os.Environ(), "TEST_SUBPROCESS_STATUS=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return // expected non-zero exit
	}
	t.Fatal("expected non-zero exit from subprocess")
}
