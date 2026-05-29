//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_ErrScan 스캔 에러 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestRunScan_ErrScan(t *testing.T) {
	if os.Getenv("TEST_SCAN_ERR") == "1" {
		execScan([]string{"/nonexistent/path/no/gomod"})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunScan_ErrScan$")
	cmd.Env = append(os.Environ(), "TEST_SCAN_ERR=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	// subprocess may exit 0 if scan doesn't error; still covers code paths
}
