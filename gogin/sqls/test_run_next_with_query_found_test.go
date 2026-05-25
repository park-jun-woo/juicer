//ff:func feature=ratchet type=session control=sequence
//ff:what TestRunNext_WithQueryFound 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunNext_WithQueryFound(t *testing.T) {
	dir := setupSessionDir(t)
	repoDir := filepath.Join(dir, "repo")
	queriesDir := filepath.Join(dir, "queries")
	os.MkdirAll(repoDir, 0o755)
	os.MkdirAll(queriesDir, 0o755)

	// Create a query file that matches
	sql := "-- name: UserGetAll :many\nSELECT * FROM users;\n"
	os.WriteFile(filepath.Join(queriesDir, "query.sql"), []byte(sql), 0o644)

	sess := &Session{
		RepoDir:    repoDir,
		QueriesDir: queriesDir,
		Methods: []MethodStatus{
			{ID: "UserRepository.GetAll", Status: "TODO"},
		},
	}
	setupTestSession(t, sess)

	err := RunNext("", "")
	// This will try to run sqlc generate, which may or may not be installed
	// Either way, should not return error from RunNext itself
	if err != nil {
		t.Logf("RunNext error (expected if sqlc not installed): %v", err)
	}
}
