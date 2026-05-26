//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_WriteErr 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestRunSQL_WriteErr(t *testing.T) {
	if os.Getenv("RSQL_ERR_WRITE") == "1" {
		dir := t.TempDir()
		runSQL([]string{"-o", "/dev/null/impossible", dir})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunSQL_WriteErr$")
	cmd.Env = append(os.Environ(), "RSQL_ERR_WRITE=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit")
}
