//ff:func feature=sql type=test control=sequence
//ff:what TestQueryExists_NotFoundCov 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestQueryExists_NotFoundCov(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "queries.sql"), []byte("-- name: OtherQuery :one\nSELECT 1"), 0644)
	if queryExists(dir, "Missing") {
		t.Fatal("expected false")
	}
}
