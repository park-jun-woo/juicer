//ff:func feature=ddl type=parse control=sequence
//ff:what TestRun 테스트
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	dir := t.TempDir()

	// Write a simple migration file
	sql := `CREATE TABLE users (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL
);
`
	if err := os.WriteFile(filepath.Join(dir, "001_init.up.sql"), []byte(sql), 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := Run(dir)
	if err != nil {
		t.Fatalf("Run() error: %v", err)
	}
	if got == "" {
		t.Fatal("Run() returned empty string")
	}
	// Should contain CREATE TABLE users
	if !contains(got, "CREATE TABLE users") {
		t.Errorf("Run() output missing CREATE TABLE users, got:\n%s", got)
	}
}
