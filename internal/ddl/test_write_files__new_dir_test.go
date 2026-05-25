//ff:func feature=ddl type=render control=sequence
//ff:what TestWriteFiles_NewDir 테스트
package ddl

import (
	"path/filepath"
	"testing"
)

func TestWriteFiles_NewDir(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "subdir")
	tables := map[string]*Table{
		"items": {Name: "items", Columns: []Column{{Name: "id", Raw: "id INT"}}},
	}
	if err := WriteFiles(tables, dir); err != nil {
		t.Fatal(err)
	}
}
