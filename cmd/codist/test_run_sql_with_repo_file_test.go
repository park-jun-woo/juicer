//ff:func feature=scan type=command control=sequence
//ff:what TestRunSQL_WithRepoFile 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunSQL_WithRepoFile(t *testing.T) {
	dir := t.TempDir()
	// Create a fake _repo.go file with SQL
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
	os.WriteFile(filepath.Join(dir, "user_repo.go"), []byte(repoGo), 0o644)

	outFile := filepath.Join(dir, "output.yaml")
	execSQL([]string{"-o", outFile, dir})
}
