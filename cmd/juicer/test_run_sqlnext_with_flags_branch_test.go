//ff:func feature=ratchet type=test control=sequence
//ff:what TestRunSQLNext_WithFlagsBranch 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunSQLNext_WithFlagsBranch(t *testing.T) {
	dir := t.TempDir()
	repoDir := filepath.Join(dir, "repo")
	queriesDir := filepath.Join(dir, "queries")
	os.MkdirAll(repoDir, 0o755)
	os.MkdirAll(queriesDir, 0o755)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	runSQLNext([]string{"-repo", repoDir, "-queries", queriesDir})
}
