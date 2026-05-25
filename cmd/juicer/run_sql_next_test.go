//ff:func feature=ratchet type=command control=sequence
//ff:what TestRunSQLNext 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunSQLNext_EmptySession(t *testing.T) {
	dir := t.TempDir()
	repoDir := filepath.Join(dir, "repo")
	queriesDir := filepath.Join(dir, "queries")
	os.MkdirAll(repoDir, 0o755)
	os.MkdirAll(queriesDir, 0o755)
	sessionDir := filepath.Join(dir, ".huma")
	os.MkdirAll(sessionDir, 0o755)
	sessionJSON := `{"repo_dir":"` + repoDir + `","queries_dir":"` + queriesDir + `","methods":[]}`
	os.WriteFile(filepath.Join(sessionDir, "sql-session.json"), []byte(sessionJSON), 0o644)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	runSQLNext([]string{})
}
