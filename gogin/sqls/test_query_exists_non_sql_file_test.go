//ff:func feature=ratchet type=session control=sequence
//ff:what TestQueryExists_NonSqlFile 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestQueryExists_NonSqlFile(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "readme.txt"), []byte("-- name: GetAll"), 0o644)
	os.MkdirAll(filepath.Join(dir, "subdir"), 0o755)

	if queryExists(dir, "GetAll") {
		t.Error("expected false for non-.sql file")
	}
}
