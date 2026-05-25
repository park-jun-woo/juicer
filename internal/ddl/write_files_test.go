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

func TestWriteFiles_NewDir(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "subdir")
	tables := map[string]*Table{
		"items": {Name: "items", Columns: []Column{{Name: "id", Raw: "id INT"}}},
	}
	if err := WriteFiles(tables, dir); err != nil {
		t.Fatal(err)
	}
}

func TestWriteFiles_Empty(t *testing.T) {
	dir := t.TempDir()
	if err := WriteFiles(map[string]*Table{}, dir); err != nil {
		t.Fatal(err)
	}
}
