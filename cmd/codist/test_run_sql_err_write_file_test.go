//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_ErrWriteFile 파일 쓰기 에러 분기 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestRunSQL_ErrWriteFile(t *testing.T) {
	if os.Getenv("TEST_SQL_WRITE") == "1" {
		dir := t.TempDir()
		runSQL([]string{"-o", "/dev/null/impossible", dir})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunSQL_ErrWriteFile$")
	cmd.Env = append(os.Environ(), "TEST_SQL_WRITE=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit")
}
