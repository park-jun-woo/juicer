//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_WriteError 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestRunSQL_WriteError(t *testing.T) {
	if os.Getenv("TEST_SUBPROCESS_SQL_WRITE") == "1" {
		dir := t.TempDir()
		execSQL([]string{"-o", "/dev/null/impossible/path", dir})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunSQL_WriteError$")
	cmd.Env = append(os.Environ(), "TEST_SUBPROCESS_SQL_WRITE=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit from subprocess")
}
