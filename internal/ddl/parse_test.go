//ff:func feature=ddl type=test control=sequence
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

	// glob error: bad pattern character
	_, err = Parse("/tmp/[")
	if err == nil {
		t.Fatal("expected glob error")
	}

	// ReadFile error: unreadable file
	dir2 := t.TempDir()
	f := filepath.Join(dir2, "001.up.sql")
	os.WriteFile(f, []byte("x"), 0o644)
	os.Chmod(f, 0o000)
	defer os.Chmod(f, 0o644)
	_, err = Parse(dir2)
	if err == nil {
		t.Fatal("expected read error")
	}
}
