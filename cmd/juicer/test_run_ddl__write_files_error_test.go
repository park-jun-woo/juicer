//ff:func feature=ddl type=test control=sequence
//ff:what TestRunDDL_WriteFilesError 테스트
package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestRunDDL_WriteFilesError(t *testing.T) {
	if os.Getenv("TEST_SUBPROCESS_DDL_WRITE") == "1" {
		dir := t.TempDir()
		sql := "CREATE TABLE users (id INT PRIMARY KEY, name TEXT);\n"
		os.WriteFile(filepath.Join(dir, "001.sql"), []byte(sql), 0o644)
		// Write to a read-only directory to trigger WriteFiles error
		roDir := t.TempDir()
		os.Chmod(roDir, 0o555)
		defer os.Chmod(roDir, 0o755)
		runDDL([]string{"-o", filepath.Join(roDir, "sub"), dir})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunDDL_WriteFilesError$")
	cmd.Env = append(os.Environ(), "TEST_SUBPROCESS_DDL_WRITE=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit from subprocess")
}
