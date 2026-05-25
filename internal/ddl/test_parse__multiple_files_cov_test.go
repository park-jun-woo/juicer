//ff:func feature=ddl type=parse control=sequence
//ff:what TestParse_MultipleFilesCov 테스트
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

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
