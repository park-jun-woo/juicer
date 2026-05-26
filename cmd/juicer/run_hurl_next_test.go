//ff:func feature=hurl type=command control=sequence
//ff:what TestRunHurlNext 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunHurlNext_WithSession(t *testing.T) {
	dir := t.TempDir()
	sessionDir := filepath.Join(dir, ".juicer")
	os.MkdirAll(sessionDir, 0o755)
	sessionJSON := `{"host":"http://localhost","tests_dir":"tests","repo_dir":"repo","endpoints":[]}`
	os.WriteFile(filepath.Join(sessionDir, "hurl-session.json"), []byte(sessionJSON), 0o644)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	runHurlNext([]string{})
}
