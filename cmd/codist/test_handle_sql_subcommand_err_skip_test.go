//ff:func feature=sql type=test control=sequence
//ff:what TestHandleSQLSubcommand_ErrSkip 에러 분기 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestHandleSQLSubcommand_ErrSkip(t *testing.T) {
	if os.Getenv("TEST_HSSC_SKIP") == "1" {
		_, cleanup := helperSetupBrokenSession(t)
		defer cleanup()
		execSQLSub([]string{"skip"})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestHandleSQLSubcommand_ErrSkip$")
	cmd.Env = append(os.Environ(), "TEST_HSSC_SKIP=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit")
}
