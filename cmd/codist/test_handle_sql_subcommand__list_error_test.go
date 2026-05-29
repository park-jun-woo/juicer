//ff:func feature=sql type=test control=sequence
//ff:what TestHandleSQLSubcommand_ListError 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

// TestHandleSQLSubcommand_ListError tests that the list error branch calls os.Exit.
func TestHandleSQLSubcommand_ListError(t *testing.T) {
	if os.Getenv("TEST_SUBPROCESS_LIST") == "1" {
		_, cleanup := helperSetupBrokenSession(t)
		defer cleanup()
		execSQLSub([]string{"list"})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestHandleSQLSubcommand_ListError$")
	cmd.Env = append(os.Environ(), "TEST_SUBPROCESS_LIST=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit from subprocess")
}
