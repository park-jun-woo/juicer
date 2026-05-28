//ff:func feature=ddl type=test control=sequence
//ff:what _down.sql 접미사 파일이 Parse에서 제외되는지 테스트
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParse_DownSQLExcluded(t *testing.T) {
	dir := t.TempDir()

	os.WriteFile(filepath.Join(dir, "001_create.up.sql"),
		[]byte("CREATE TABLE orders (id INT PRIMARY KEY);"), 0o644)

	// _down.sql variant
	os.WriteFile(filepath.Join(dir, "001_create_down.sql"),
		[]byte("DROP TABLE orders;"), 0o644)

	tables, err := Parse(dir)
	if err != nil {
		t.Fatal(err)
	}
	if tables["orders"] == nil {
		t.Fatal("expected orders table — down migration should have been excluded")
	}
}
