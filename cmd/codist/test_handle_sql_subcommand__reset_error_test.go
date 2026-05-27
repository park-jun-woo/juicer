//ff:func feature=sql type=test control=sequence
//ff:what TestHandleSQLSubcommand_ResetError 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

// TestHandleSQLSubcommand_ResetError tests that the reset error branch calls os.Exit.
func TestHandleSQLSubcommand_ResetError(t *testing.T) {
	if os.Getenv("TEST_SUBPROCESS_RESET") == "1" {
		_, cleanup := helperSetupUndeletableSession(t)
		defer cleanup()
		handleSQLSubcommand([]string{"reset"})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestHandleSQLSubcommand_ResetError$")
	cmd.Env = append(os.Environ(), "TEST_SUBPROCESS_RESET=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit from subprocess")
}
