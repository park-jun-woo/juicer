//ff:func feature=ddl type=parse control=sequence
//ff:what TestParse 테스트
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	dir := t.TempDir()

	// Migration 1: create users table
	sql1 := `CREATE TABLE users (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL
);
`
	// Migration 2: add email column
	sql2 := `ALTER TABLE users ADD COLUMN email TEXT;
`
	if err := os.WriteFile(filepath.Join(dir, "001_init.up.sql"), []byte(sql1), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "002_add_email.up.sql"), []byte(sql2), 0o644); err != nil {
		t.Fatal(err)
	}

	tables, err := Parse(dir)
	if err != nil {
		t.Fatalf("Parse() error: %v", err)
	}
	if len(tables) != 1 {
		t.Fatalf("expected 1 table, got %d", len(tables))
	}
	tbl, ok := tables["users"]
	if !ok {
		t.Fatal("table 'users' not found")
	}
	if len(tbl.Columns) != 3 {
		t.Errorf("expected 3 columns, got %d", len(tbl.Columns))
	}
}
