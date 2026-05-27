//ff:func feature=sql type=test control=sequence
//ff:what helperSetupUndeletableSession 헬퍼 함수
package main

import (
	"os"
	"path/filepath"
	"testing"
)

// helperSetupUndeletableSession creates a session file in a read-only directory
// so DeleteSession returns an error. Returns dir path and cleanup func.
func helperSetupUndeletableSession(t *testing.T) (string, func()) {
	t.Helper()
	dir := t.TempDir()
	sessionDir := filepath.Join(dir, ".codist")
	os.MkdirAll(sessionDir, 0o755)
	os.WriteFile(filepath.Join(sessionDir, "sql-session.json"), []byte(`{"repo_dir":".","queries_dir":".","methods":[]}`), 0o644)
	// Make the directory read-only so file removal fails
	os.Chmod(sessionDir, 0o555)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	return dir, func() {
		os.Chmod(sessionDir, 0o755) // restore so t.TempDir cleanup works
		os.Chdir(oldWd)
	}
}
