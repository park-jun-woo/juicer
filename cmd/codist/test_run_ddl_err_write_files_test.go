//ff:func feature=ddl type=test control=sequence
//ff:what TestRunDDL_ErrWriteFiles 파일 쓰기 에러 분기 테스트
package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestRunDDL_ErrWriteFiles(t *testing.T) {
	if os.Getenv("TEST_DDL_WRITE_ERR") == "1" {
		dir := t.TempDir()
		sql := "CREATE TABLE users (id INT PRIMARY KEY, name TEXT);\n"
		os.WriteFile(filepath.Join(dir, "001.sql"), []byte(sql), 0o644)
		roDir := t.TempDir()
		os.Chmod(roDir, 0o555)
		defer os.Chmod(roDir, 0o755)
		execDDL([]string{"-o", filepath.Join(roDir, "sub"), dir})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunDDL_ErrWriteFiles$")
	cmd.Env = append(os.Environ(), "TEST_DDL_WRITE_ERR=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit")
}
