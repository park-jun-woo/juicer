//ff:func feature=ddl type=test control=sequence
//ff:what TestWriteFiles_NewDirCov 테스트
package ddl

import (
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
	matches, _ := filepath.Glob(filepath.Join(outDir, "*_orders.sql"))
	if len(matches) == 0 {
		t.Fatalf("expected a *_orders.sql file in %s", outDir)
	}
}
