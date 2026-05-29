//ff:func feature=ddl type=test control=sequence
//ff:what TestWriteFiles_MkdirErr 테스트
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFiles_MkdirErr(t *testing.T) {
	// Try writing to a path where a file blocks dir creation
	dir := t.TempDir()
	blocker := filepath.Join(dir, "block")
	os.WriteFile(blocker, []byte("x"), 0o644)
	tables := map[string]*Table{"a": {Name: "a"}}
	err := WriteFiles(nil, tables, filepath.Join(blocker, "sub"))
	if err == nil {
		t.Fatal("expected error")
	}
}
