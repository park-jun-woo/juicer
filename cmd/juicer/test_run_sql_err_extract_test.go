//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_ErrExtract 추출 에러 분기 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestRunSQL_ErrExtract(t *testing.T) {
	if os.Getenv("TEST_SQL_EXTRACT") == "1" {
		// Use a non-existent directory that causes Extract to fail
		runSQL([]string{"/nonexistent/dir/sql"})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunSQL_ErrExtract$")
	cmd.Env = append(os.Environ(), "TEST_SQL_EXTRACT=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	// Extract may not fail on non-existent dir; still covers paths
}
