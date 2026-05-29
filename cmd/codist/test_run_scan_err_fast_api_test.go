//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_FastAPI FastAPI 스캐너 실행 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestRunScan_FastAPI(t *testing.T) {
	if os.Getenv("TEST_SCAN_FASTAPI") == "1" {
		dir := setupMinimalGoProject(t)
		execScan([]string{"--framework", "fastapi", dir})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunScan_FastAPI$")
	cmd.Env = append(os.Environ(), "TEST_SCAN_FASTAPI=1")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("expected zero exit, got %v", err)
	}
}
