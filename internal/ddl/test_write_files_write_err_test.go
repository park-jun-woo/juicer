//ff:func feature=ddl type=test control=sequence
//ff:what TestWriteFiles_WriteErr 테스트
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFiles_WriteErr(t *testing.T) {
	dir := t.TempDir()
	outDir := filepath.Join(dir, "out")
	os.MkdirAll(outDir, 0o755)
	os.Chmod(outDir, 0o444)
	defer os.Chmod(outDir, 0o755)
	tables := map[string]*Table{"a": {Name: "a", Columns: []Column{{Name: "id", Raw: "id INT"}}}}
	err := WriteFiles(nil, tables, outDir)
	if err == nil {
		t.Fatal("expected error")
	}
}
