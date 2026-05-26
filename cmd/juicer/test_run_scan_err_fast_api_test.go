//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_ErrFastAPI FastAPI 에러 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestRunScan_ErrFastAPI(t *testing.T) {
	if os.Getenv("TEST_SCAN_FASTAPI") == "1" {
		dir := setupMinimalGoProject(t)
		runScan([]string{"-framework", "fastapi", dir})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunScan_ErrFastAPI$")
	cmd.Env = append(os.Environ(), "TEST_SCAN_FASTAPI=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit")
}
