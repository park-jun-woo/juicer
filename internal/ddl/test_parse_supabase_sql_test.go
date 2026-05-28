//ff:func feature=ddl type=test control=sequence
//ff:what Parse가 Supabase *.sql 파일과 기존 *.up.sql 모두 인식하는지 테스트
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParse_SupabaseSQL(t *testing.T) {
	dir := t.TempDir()

	// Supabase-style migration file (no .up.sql suffix)
	os.WriteFile(filepath.Join(dir, "20240101000000_create_profiles.sql"),
		[]byte("CREATE TABLE profiles (id UUID PRIMARY KEY, name TEXT);"), 0o644)

	// Traditional .up.sql file
	os.WriteFile(filepath.Join(dir, "001_init.up.sql"),
		[]byte("CREATE TABLE users (id INT PRIMARY KEY);"), 0o644)

	// Down migration file — should be excluded
	os.WriteFile(filepath.Join(dir, "001_init.down.sql"),
		[]byte("DROP TABLE users;"), 0o644)

	tables, err := Parse(dir)
	if err != nil {
		t.Fatal(err)
	}

	if tables["users"] == nil {
		t.Fatal("expected users table from .up.sql")
	}
	if tables["profiles"] == nil {
		t.Fatal("expected profiles table from Supabase .sql")
	}
}
