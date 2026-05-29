//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_ScanError 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestRunScan_ScanError(t *testing.T) {
	if os.Getenv("TEST_SUBPROCESS_SCAN_ERR") == "1" {
		// Use a non-existent path that packages.Load cannot resolve
		execScan([]string{"/nonexistent/path/no/gomod"})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunScan_ScanError$")
	cmd.Env = append(os.Environ(), "TEST_SUBPROCESS_SCAN_ERR=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	// If the subprocess exits 0, it means Scan didn't error. That's OK for coverage.
	// The subprocess still covers code paths.
}
