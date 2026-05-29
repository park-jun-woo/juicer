//ff:func feature=ratchet type=test control=sequence
//ff:what runSQLNext 비-Exit 분기 테스트 (os.Exit 에러 분기는 서브프로세스 테스트로 분리)
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunSQLNext_FlagsBranch(t *testing.T) {
	dir := t.TempDir()
	repoDir := filepath.Join(dir, "repo")
	queriesDir := filepath.Join(dir, "queries")
	os.MkdirAll(repoDir, 0o755)
	os.MkdirAll(queriesDir, 0o755)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	// success path; the err != nil branch on line 19 calls os.Exit and
	// cannot be covered in-process.
	execSQLNext([]string{"--repo", repoDir, "--queries", queriesDir})
}
