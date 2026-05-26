//ff:func feature=sql type=test control=sequence
//ff:what TestQueryExists_NonSQLFileCov 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestQueryExists_NonSQLFileCov(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "readme.txt"), []byte("-- name: X :one"), 0644)
	if queryExists(dir, "X") {
		t.Fatal("expected false for non-sql file")
	}
}
