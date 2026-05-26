//ff:func feature=sql type=test control=sequence
//ff:what TestQueryExists_Found 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestQueryExists_Found(t *testing.T) {
	dir, err := os.MkdirTemp("", "qe-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	os.WriteFile(filepath.Join(dir, "queries.sql"), []byte("-- name: FindUser :one\nSELECT * FROM users"), 0644)
	if !queryExists(dir, "FindUser") {
		t.Fatal("expected true")
	}
}
