//ff:func feature=ratchet type=test control=sequence
//ff:what TestRunSQLNext_ErrBranch 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestRunSQLNext_ErrBranch(t *testing.T) {
	if os.Getenv("RSN_ERR") == "1" {
		_, cleanup := helperSetupBrokenSession(t)
		defer cleanup()
		runSQLNext([]string{})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunSQLNext_ErrBranch$")
	cmd.Env = append(os.Environ(), "RSN_ERR=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit")
}
