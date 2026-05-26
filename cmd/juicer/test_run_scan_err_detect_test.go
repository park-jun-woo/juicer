//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_ErrDetect 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestRunScan_ErrDetect(t *testing.T) {
	if os.Getenv("RS_ERR_DETECT") == "1" {
		dir := t.TempDir()
		runScan([]string{dir})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunScan_ErrDetect$")
	cmd.Env = append(os.Environ(), "RS_ERR_DETECT=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit")
}
