//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_ErrNestJS NestJS 에러 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestRunScan_ErrNestJS(t *testing.T) {
	if os.Getenv("TEST_SCAN_NESTJS") == "1" {
		dir := t.TempDir()
		runScan([]string{"-framework", "nestjs", dir})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunScan_ErrNestJS$")
	cmd.Env = append(os.Environ(), "TEST_SCAN_NESTJS=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit")
}
