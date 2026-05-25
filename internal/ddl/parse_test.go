//ff:func feature=ddl type=parse control=sequence
//ff:what TestParse_Basic 테스트
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParse_Basic(t *testing.T) {
	dir := t.TempDir()
	sql := "CREATE TABLE users (id INT PRIMARY KEY, name TEXT);\n"
	os.WriteFile(filepath.Join(dir, "001_init.up.sql"), []byte(sql), 0o644)
	tables, err := Parse(dir)
	if err != nil {
		t.Fatal(err)
	}
	if tables["users"] == nil {
		t.Fatal("expected users table")
	}
}
