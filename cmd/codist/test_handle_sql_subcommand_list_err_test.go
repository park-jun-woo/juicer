//ff:func feature=sql type=test control=sequence
//ff:what TestHandleSQLSubcommand_ListErr 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestHandleSQLSubcommand_ListErr(t *testing.T) {
	if os.Getenv("HSSC_ERR_LIST") == "1" {
		_, cleanup := helperSetupBrokenSession(t)
		defer cleanup()
		execSQLSub([]string{"list"})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestHandleSQLSubcommand_ListErr$")
	cmd.Env = append(os.Environ(), "HSSC_ERR_LIST=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit")
}
