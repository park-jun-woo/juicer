//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_FastAPIBranch FastAPI 스캐너 실행 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestRunScan_FastAPIBranch(t *testing.T) {
	if os.Getenv("RS_ERR_FASTAPI") == "1" {
		dir := setupMinimalGoProject(t)
		runScan([]string{"-framework", "fastapi", dir})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunScan_FastAPIBranch$")
	cmd.Env = append(os.Environ(), "RS_ERR_FASTAPI=1")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("expected zero exit, got %v", err)
	}
}
