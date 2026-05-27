//ff:func feature=sql type=test control=sequence
//ff:what TestHandleSQLSubcommand_ErrStatus 에러 분기 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestHandleSQLSubcommand_ErrStatus(t *testing.T) {
	if os.Getenv("TEST_HSSC_STATUS") == "1" {
		_, cleanup := helperSetupBrokenSession(t)
		defer cleanup()
		handleSQLSubcommand([]string{"status"})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestHandleSQLSubcommand_ErrStatus$")
	cmd.Env = append(os.Environ(), "TEST_HSSC_STATUS=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit")
}
