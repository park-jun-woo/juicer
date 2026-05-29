//ff:func feature=ratchet type=test control=sequence
//ff:what TestRunSQLNext_ErrRunNext 에러 분기 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestRunSQLNext_ErrRunNext(t *testing.T) {
	if os.Getenv("TEST_SQLNEXT_ERR") == "1" {
		_, cleanup := helperSetupBrokenSession(t)
		defer cleanup()
		execSQLNext([]string{})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunSQLNext_ErrRunNext$")
	cmd.Env = append(os.Environ(), "TEST_SQLNEXT_ERR=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit")
}
