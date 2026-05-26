//ff:func feature=ratchet type=session control=sequence
//ff:what TestPrintSkeleton_WithDynamicAndParams 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPrintSkeleton_WithDynamicAndParams(t *testing.T) {
	dir := setupSessionDir(t)
	repoDir := filepath.Join(dir, "repo")
	os.MkdirAll(repoDir, 0o755)

	// Create a repo file with params and dynamic SQL
	repoGo := `package repo

import "context"
import "database/sql"
import "fmt"

type UserRepo struct {
	db *sql.DB
}

func (r *UserRepo) Search(ctx context.Context, name string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT id, name FROM users WHERE name = '%s'", name)
	return r.db.QueryContext(ctx, query)
}
`
	os.WriteFile(filepath.Join(repoDir, "user_repo.go"), []byte(repoGo), 0o644)

	sess := &Session{
		RepoDir:    repoDir,
		QueriesDir: ".",
		Methods: []MethodStatus{
			{ID: "UserRepo.Search", Status: "TODO"},
		},
	}
	result, _ := Extract(repoDir)
	var methods []MethodSkeleton
	if result != nil {
		methods = result.Methods
	}
	printSkeleton(sess, 0, methods)
}
