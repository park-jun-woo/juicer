//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_ErrWriteFile 파일 쓰기 에러 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestRunScan_ErrWriteFile(t *testing.T) {
	if os.Getenv("TEST_SCAN_WRITE") == "1" {
		dir := setupMinimalGoProject(t)
		execScan([]string{"-o", "/dev/null/impossible/path", dir})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunScan_ErrWriteFile$")
	cmd.Env = append(os.Environ(), "TEST_SCAN_WRITE=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit")
}
