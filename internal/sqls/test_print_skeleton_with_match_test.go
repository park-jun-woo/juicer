//ff:func feature=ratchet type=session control=sequence
//ff:what TestPrintSkeleton_WithMatch 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPrintSkeleton_WithMatch(t *testing.T) {
	dir := setupSessionDir(t)
	repoDir := filepath.Join(dir, "repo")
	os.MkdirAll(repoDir, 0o755)

	// Create a repo file
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

	sess := &Session{
		RepoDir:    repoDir,
		QueriesDir: ".",
		Methods: []MethodStatus{
			{ID: "UserRepo.GetAll", Status: "TODO"},
		},
	}
	result, _ := Extract(repoDir)
	var methods []MethodSkeleton
	if result != nil {
		methods = result.Methods
	}
	printSkeleton(sess, 0, methods)
}
