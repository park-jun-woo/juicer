//ff:func feature=ratchet type=test control=sequence
//ff:what TestRunSQLNext_RunNextError 테스트
package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestRunSQLNext_RunNextError(t *testing.T) {
	if os.Getenv("TEST_SUBPROCESS_SQLNEXT_ERR") == "1" {
		dir := t.TempDir()
		sessionDir := filepath.Join(dir, ".codist")
		os.MkdirAll(sessionDir, 0o755)
		os.WriteFile(filepath.Join(sessionDir, "sql-session.json"), []byte("INVALID"), 0o644)
		oldWd, _ := os.Getwd()
		os.Chdir(dir)
		defer os.Chdir(oldWd)
		runSQLNext([]string{})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunSQLNext_RunNextError$")
	cmd.Env = append(os.Environ(), "TEST_SUBPROCESS_SQLNEXT_ERR=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit from subprocess")
}
