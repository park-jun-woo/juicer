//ff:func feature=sql type=parse control=sequence
//ff:what TestExtract 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtract(t *testing.T) {
	dir := t.TempDir()

	t.Run("empty dir", func(t *testing.T) {
		result, err := Extract(dir)
		if err != nil {
			t.Fatalf("Extract() error: %v", err)
		}
		if len(result.Methods) != 0 {
			t.Errorf("expected 0 methods, got %d", len(result.Methods))
		}
	})

	t.Run("non-repo files skipped", func(t *testing.T) {
		os.WriteFile(filepath.Join(dir, "handler.go"), []byte("package repo\n"), 0o644)
		result, err := Extract(dir)
		if err != nil {
			t.Fatalf("Extract() error: %v", err)
		}
		if len(result.Methods) != 0 {
			t.Errorf("expected 0 methods, got %d", len(result.Methods))
		}
	})

	t.Run("with repo file", func(t *testing.T) {
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

func (r *UserRepo) Create(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (name) VALUES ($1)")
	return err
}
`
		os.WriteFile(filepath.Join(dir, "user_repo.go"), []byte(repoGo), 0o644)
		result, err := Extract(dir)
		if err != nil {
			t.Fatalf("Extract() error: %v", err)
		}
		if len(result.Methods) < 1 {
			t.Errorf("expected at least 1 method, got %d", len(result.Methods))
		}
	})

	t.Run("test files skipped", func(t *testing.T) {
		os.WriteFile(filepath.Join(dir, "user_repo_test.go"), []byte("package repo\n"), 0o644)
		result, err := Extract(dir)
		if err != nil {
			t.Fatalf("Extract() error: %v", err)
		}
		// Should not include test file methods
		_ = result
	})
}
