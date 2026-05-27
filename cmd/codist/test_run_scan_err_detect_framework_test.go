//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_ErrDetectFramework 프레임워크 감지 에러 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestRunScan_ErrDetectFramework(t *testing.T) {
	if os.Getenv("TEST_SCAN_DETECT") == "1" {
		dir := t.TempDir()
		runScan([]string{dir})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunScan_ErrDetectFramework$")
	cmd.Env = append(os.Environ(), "TEST_SCAN_DETECT=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit")
}
