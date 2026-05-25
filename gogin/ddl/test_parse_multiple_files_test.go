//ff:func feature=ddl type=parse control=sequence
//ff:what TestParse_MultipleFiles 테스트
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParse_MultipleFiles(t *testing.T) {
	dir := t.TempDir()

	sql1 := `CREATE TABLE users (id BIGINT PRIMARY KEY);`
	sql2 := `CREATE TABLE posts (id BIGINT PRIMARY KEY, user_id BIGINT);`

	os.WriteFile(filepath.Join(dir, "001_users.up.sql"), []byte(sql1), 0o644)
	os.WriteFile(filepath.Join(dir, "002_posts.up.sql"), []byte(sql2), 0o644)

	tables, err := Parse(dir)
	if err != nil {
		t.Fatalf("Parse() error: %v", err)
	}
	if len(tables) != 2 {
		t.Fatalf("expected 2 tables, got %d", len(tables))
	}
	if _, ok := tables["users"]; !ok {
		t.Error("table 'users' not found")
	}
	if _, ok := tables["posts"]; !ok {
		t.Error("table 'posts' not found")
	}
}
