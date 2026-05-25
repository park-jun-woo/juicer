//ff:func feature=ratchet type=session control=sequence
//ff:what TestQueryExists 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestQueryExists(t *testing.T) {
	dir := t.TempDir()

	// Create a .sql file with query
	sql := `-- name: GetAll :many
SELECT * FROM users;
`
	os.WriteFile(filepath.Join(dir, "queries.sql"), []byte(sql), 0o644)

	if !queryExists(dir, "GetAll") {
		t.Error("expected query 'GetAll' to exist")
	}
	if queryExists(dir, "NonExistent") {
		t.Error("expected query 'NonExistent' to not exist")
	}
}
