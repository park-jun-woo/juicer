//ff:func feature=ddl type=render control=sequence
//ff:what TestWriteFiles_Basic 테스트
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFiles_Basic(t *testing.T) {
	dir := t.TempDir()
	tables := map[string]*Table{
		"users": {Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}}},
	}
	if err := WriteFiles(tables, dir); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(dir, "users.sql")); err != nil {
		t.Fatalf("expected users.sql: %v", err)
	}
}
