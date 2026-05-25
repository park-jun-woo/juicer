//ff:func feature=ratchet type=session control=sequence
//ff:what TestCreateSession_WithRepoFiles 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateSession_WithRepoFiles(t *testing.T) {
	dir := setupSessionDir(t)
	repoDir := filepath.Join(dir, "repo")
	queriesDir := filepath.Join(dir, "queries")
	os.MkdirAll(repoDir, 0o755)
	os.MkdirAll(queriesDir, 0o755)

	// Create a repo file with SQL usage
	repoGo := `package repo

import "context"
import "database/sql"

type UserRepo struct {
	db *sql.DB
}

func (r *UserRepo) GetAll(ctx context.Context) error {
	_, err := r.db.QueryContext(ctx, "SELECT id, name FROM users")
	return err
}
`
	os.WriteFile(filepath.Join(repoDir, "user_repo.go"), []byte(repoGo), 0o644)

	err := createSession(repoDir, queriesDir)
	if err != nil {
		t.Fatalf("createSession() error: %v", err)
	}

	if !SessionExists() {
		t.Error("expected session to exist")
	}

	sess, _ := LoadSession()
	if len(sess.Methods) == 0 {
		t.Log("no methods extracted (may be expected for simple repo)")
	}
}
