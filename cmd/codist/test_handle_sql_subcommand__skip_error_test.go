//ff:func feature=sql type=test control=sequence
//ff:what TestHandleSQLSubcommand_SkipError 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

// TestHandleSQLSubcommand_SkipError tests that the skip error branch calls os.Exit.
func TestHandleSQLSubcommand_SkipError(t *testing.T) {
	if os.Getenv("TEST_SUBPROCESS_SKIP") == "1" {
		_, cleanup := helperSetupBrokenSession(t)
		defer cleanup()
		execSQLSub([]string{"skip"})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestHandleSQLSubcommand_SkipError$")
	cmd.Env = append(os.Environ(), "TEST_SUBPROCESS_SKIP=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit from subprocess")
}
