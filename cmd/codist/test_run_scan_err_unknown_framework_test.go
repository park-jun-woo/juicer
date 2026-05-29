//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_ErrUnknownFramework 알 수 없는 프레임워크 에러 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestRunScan_ErrUnknownFramework(t *testing.T) {
	if os.Getenv("TEST_SCAN_UNKNOWN_FW") == "1" {
		dir := setupMinimalGoProject(t)
		execScan([]string{"--framework", "unknown", dir})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunScan_ErrUnknownFramework$")
	cmd.Env = append(os.Environ(), "TEST_SCAN_UNKNOWN_FW=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit")
}
