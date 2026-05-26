//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_NestJSEmpty NestJS 빈 프로젝트 스캔 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestRunScan_NestJSEmpty(t *testing.T) {
	if os.Getenv("TEST_SCAN_NESTJS") == "1" {
		dir := t.TempDir()
		runScan([]string{"-framework", "nestjs", dir})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunScan_NestJSEmpty$")
	cmd.Env = append(os.Environ(), "TEST_SCAN_NESTJS=1")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("expected zero exit, got %v", err)
	}
}
