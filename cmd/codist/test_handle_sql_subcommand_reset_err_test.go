//ff:func feature=sql type=test control=sequence
//ff:what TestHandleSQLSubcommand_ResetErr 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestHandleSQLSubcommand_ResetErr(t *testing.T) {
	if os.Getenv("HSSC_ERR_RESET") == "1" {
		_, cleanup := helperSetupUndeletableSession(t)
		defer cleanup()
		execSQLSub([]string{"reset"})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestHandleSQLSubcommand_ResetErr$")
	cmd.Env = append(os.Environ(), "HSSC_ERR_RESET=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit")
}
