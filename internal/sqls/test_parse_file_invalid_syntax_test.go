//ff:func feature=sql type=test control=sequence
//ff:what TestParseFile_InvalidSyntax 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFile_InvalidSyntax(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "bad.go")
	os.WriteFile(f, []byte("invalid go code!!!"), 0o644)
	_, err := parseFile(f)
	if err == nil {
		t.Fatal("expected error")
	}
}
