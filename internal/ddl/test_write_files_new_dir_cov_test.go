//ff:func feature=ddl type=test control=sequence
//ff:what TestWriteFiles_NewDirCov 테스트
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFiles_NewDirCov(t *testing.T) {
	dir := t.TempDir()
	outDir := filepath.Join(dir, "new", "sub")
	tables := map[string]*Table{
		"orders": {Name: "orders", Columns: []Column{{Name: "id", Raw: "id INT"}}},
	}
	if err := WriteFiles(nil, tables, outDir); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(outDir, "orders.sql")); err != nil {
		t.Fatalf("expected orders.sql: %v", err)
	}
}
