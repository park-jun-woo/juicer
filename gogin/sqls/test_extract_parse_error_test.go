//ff:func feature=sql type=parse control=sequence
//ff:what TestExtract_ParseError 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtract_ParseError(t *testing.T) {
	dir := t.TempDir()
	// Write invalid Go code
	os.WriteFile(filepath.Join(dir, "bad_repo.go"), []byte("not valid go code{{{"), 0o644)
	_, err := Extract(dir)
	if err == nil {
		t.Error("expected error for invalid Go code")
	}
}
