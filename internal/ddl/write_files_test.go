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
	if err := WriteFiles(nil, tables, dir); err != nil {
		t.Fatal(err)
	}
	matches, _ := filepath.Glob(filepath.Join(dir, "*_users.sql"))
	if len(matches) == 0 {
		t.Fatalf("expected a *_users.sql file in %s", dir)
	}

	// MkdirAll error
	err := WriteFiles(nil, tables, "/dev/null/impossible")
	if err == nil {
		t.Fatal("expected mkdir error")
	}

	// WriteFile error (read-only directory)
	roDir := t.TempDir()
	os.Chmod(roDir, 0o555)
	defer os.Chmod(roDir, 0o755)
	err = WriteFiles(nil, tables, filepath.Join(roDir, "sub"))
	if err == nil {
		t.Fatal("expected write error")
	}

	// WriteFile error — directory exists but is read-only, new file
	roDir2 := t.TempDir()
	os.Chmod(roDir2, 0o555)
	defer os.Chmod(roDir2, 0o755)
	err = WriteFiles(nil, tables, roDir2)
	if err == nil {
		t.Fatal("expected WriteFile error on read-only dir")
	}
}
