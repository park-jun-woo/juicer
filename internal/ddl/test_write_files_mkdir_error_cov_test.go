//ff:func feature=ddl type=test control=sequence
//ff:what TestWriteFiles_MkdirErrorCov 테스트
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFiles_MkdirErrorCov(t *testing.T) {
	dir := t.TempDir()
	blockPath := filepath.Join(dir, "block")
	os.WriteFile(blockPath, []byte("x"), 0o644)
	tables := map[string]*Table{
		"users": {Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}}},
	}
	err := WriteFiles(tables, filepath.Join(blockPath, "sub"))
	if err == nil {
		t.Fatal("expected error creating dir under file")
	}
}
