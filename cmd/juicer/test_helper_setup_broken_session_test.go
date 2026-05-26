//ff:func feature=sql type=test control=sequence
//ff:what helperSetupBrokenSession 헬퍼 함수
package main

import (
	"os"
	"path/filepath"
	"testing"
)

// helperSetupBrokenSession creates a directory with an invalid session file
// so LoadSession returns an error. Returns dir path and cleanup func.
func helperSetupBrokenSession(t *testing.T) (string, func()) {
	t.Helper()
	dir := t.TempDir()
	sessionDir := filepath.Join(dir, ".juicer")
	os.MkdirAll(sessionDir, 0o755)
	os.WriteFile(filepath.Join(sessionDir, "sql-session.json"), []byte("INVALID"), 0o644)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	return dir, func() { os.Chdir(oldWd) }
}
