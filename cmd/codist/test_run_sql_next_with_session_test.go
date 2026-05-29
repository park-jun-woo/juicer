//ff:func feature=scan type=command control=sequence
//ff:what TestRunSQLNext_WithSession 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunSQLNext_WithSession(t *testing.T) {
	// Create a temp directory with a pre-existing session so RunNext doesn't need --repo/--queries
	dir := t.TempDir()
	repoDir := filepath.Join(dir, "repo")
	queriesDir := filepath.Join(dir, "queries")
	os.MkdirAll(repoDir, 0o755)
	os.MkdirAll(queriesDir, 0o755)

	// Create session file
	sessionDir := filepath.Join(dir, ".codist")
	os.MkdirAll(sessionDir, 0o755)
	sessionJSON := `{
  "repo_dir": "` + repoDir + `",
  "queries_dir": "` + queriesDir + `",
  "methods": []
}`
	os.WriteFile(filepath.Join(sessionDir, "sql-session.json"), []byte(sessionJSON), 0o644)

	// Change to the dir so SessionExists() finds .codist/
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	// RunNext with empty methods should print "All queries complete!"
	execSQLNext([]string{})
}
