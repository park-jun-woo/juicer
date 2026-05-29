//ff:func feature=sql type=test control=sequence
//ff:what TestHandleSQLSubcommand_ErrReset 에러 분기 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestHandleSQLSubcommand_ErrReset(t *testing.T) {
	if os.Getenv("TEST_HSSC_RESET") == "1" {
		_, cleanup := helperSetupUndeletableSession(t)
		defer cleanup()
		execSQLSub([]string{"reset"})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestHandleSQLSubcommand_ErrReset$")
	cmd.Env = append(os.Environ(), "TEST_HSSC_RESET=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit")
}
