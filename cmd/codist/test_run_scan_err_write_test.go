//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_ErrWrite 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestRunScan_ErrWrite(t *testing.T) {
	if os.Getenv("RS_ERR_WRITE") == "1" {
		dir := setupMinimalGoProject(t)
		runScan([]string{"-o", "/dev/null/impossible/path", dir})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunScan_ErrWrite$")
	cmd.Env = append(os.Environ(), "RS_ERR_WRITE=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit")
}
