//ff:func feature=ratchet type=session control=sequence
//ff:what TestRunNext_WithTODO_NoQuery 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunNext_WithTODO_NoQuery(t *testing.T) {
	dir := setupSessionDir(t)
	repoDir := filepath.Join(dir, "repo")
	queriesDir := filepath.Join(dir, "queries")
	os.MkdirAll(repoDir, 0o755)
	os.MkdirAll(queriesDir, 0o755)

	sess := &Session{
		RepoDir:    repoDir,
		QueriesDir: queriesDir,
		Methods: []MethodStatus{
			{ID: "UserRepo.GetAll", Status: "TODO"},
		},
	}
	setupTestSession(t, sess)

	err := RunNext("", "")
	if err != nil {
		t.Fatalf("RunNext() error: %v", err)
	}
}
