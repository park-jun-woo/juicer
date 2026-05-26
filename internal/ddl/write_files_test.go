//ff:func feature=ddl type=test control=sequence
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

	// MkdirAll error
	err := WriteFiles(tables, "/dev/null/impossible")
	if err == nil {
		t.Fatal("expected mkdir error")
	}

	// WriteFile error (read-only directory)
	roDir := t.TempDir()
	os.Chmod(roDir, 0o555)
	defer os.Chmod(roDir, 0o755)
	err = WriteFiles(tables, filepath.Join(roDir, "sub"))
	if err == nil {
		t.Fatal("expected write error")
	}
}
