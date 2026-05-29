//ff:func feature=ddl type=test control=sequence
//ff:what TestRunDDL_WriteErr 테스트
package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestRunDDL_WriteErr(t *testing.T) {
	if os.Getenv("RD_ERR_WRITE") == "1" {
		dir := t.TempDir()
		sql := "CREATE TABLE users (id INT PRIMARY KEY, name TEXT);\n"
		os.WriteFile(filepath.Join(dir, "001.sql"), []byte(sql), 0o644)
		roDir := t.TempDir()
		os.Chmod(roDir, 0o555)
		defer os.Chmod(roDir, 0o755)
		execDDL([]string{"-o", filepath.Join(roDir, "sub"), dir})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunDDL_WriteErr$")
	cmd.Env = append(os.Environ(), "RD_ERR_WRITE=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit")
}
