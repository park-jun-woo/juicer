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

func TestParse_EmptyDirCov(t *testing.T) {
	dir := t.TempDir()
	tables, err := Parse(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(tables) != 0 {
		t.Fatalf("expected 0 tables, got %d", len(tables))
	}
}

func TestParse_MultipleFilesCov(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "001_init.up.sql"), []byte("CREATE TABLE users (id INT);\n"), 0o644)
	os.WriteFile(filepath.Join(dir, "002_items.up.sql"), []byte("CREATE TABLE items (id INT);\n"), 0o644)
	tables, err := Parse(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(tables) != 2 {
		t.Fatalf("expected 2 tables, got %d", len(tables))
	}
}
