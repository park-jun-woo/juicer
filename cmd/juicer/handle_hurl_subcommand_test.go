//ff:func feature=hurl type=command control=sequence
//ff:what TestHandleHurlSubcommand 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func setupHurlSession(t *testing.T) (string, func()) {
	t.Helper()
	dir := t.TempDir()
	sessionDir := filepath.Join(dir, ".juicer")
	os.MkdirAll(sessionDir, 0o755)
	sessionJSON := `{"host":"http://localhost","tests_dir":"tests","repo_dir":"repo","endpoints":[]}`
	os.WriteFile(filepath.Join(sessionDir, "hurl-session.json"), []byte(sessionJSON), 0o644)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	return dir, func() { os.Chdir(oldWd) }
}
