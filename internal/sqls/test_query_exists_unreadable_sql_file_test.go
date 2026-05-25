//ff:func feature=ratchet type=session control=sequence
//ff:what TestQueryExists_UnreadableSqlFile 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestQueryExists_UnreadableSqlFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "query.sql")
	os.WriteFile(path, []byte("-- name: GetAll :many\n"), 0o644)
	os.Chmod(path, 0o000)
	t.Cleanup(func() { os.Chmod(path, 0o644) })

	// Should not crash, returns false
	if queryExists(dir, "GetAll") {
		t.Error("expected false for unreadable file")
	}
}
