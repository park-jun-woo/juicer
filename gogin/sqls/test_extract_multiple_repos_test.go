//ff:func feature=sql type=parse control=sequence
//ff:what TestExtract_MultipleRepos 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtract_MultipleRepos(t *testing.T) {
	dir := t.TempDir()

	repoA := `package repo

import "context"
import "database/sql"

type BRepo struct {
	db *sql.DB
}

func (r *BRepo) Get(ctx context.Context) error {
	_, err := r.db.QueryContext(ctx, "SELECT id FROM b_table")
	return err
}
`
	repoB := `package repo

import "context"
import "database/sql"

type ARepo struct {
	db *sql.DB
}

func (r *ARepo) Get(ctx context.Context) error {
	_, err := r.db.QueryContext(ctx, "SELECT id FROM a_table")
	return err
}
`
	os.WriteFile(filepath.Join(dir, "b_repo.go"), []byte(repoA), 0o644)
	os.WriteFile(filepath.Join(dir, "a_repo.go"), []byte(repoB), 0o644)

	result, err := Extract(dir)
	if err != nil {
		t.Fatalf("Extract() error: %v", err)
	}
	if len(result.Methods) >= 2 {
		// Should be sorted: ARepo before BRepo
		if result.Methods[0].Repo > result.Methods[1].Repo {
			t.Error("expected methods to be sorted by Repo")
		}
	}
}
