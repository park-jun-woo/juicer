//ff:func feature=ratchet type=session control=sequence
//ff:what TestRunList_LoadError 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunList_LoadError(t *testing.T) {
	dir := setupSessionDir(t)
	os.MkdirAll(filepath.Join(dir, ".juicer"), 0o755)
	os.WriteFile(filepath.Join(dir, ".juicer", "sql-session.json"), []byte("bad"), 0o644)

	err := RunList()
	if err == nil {
		t.Error("expected error for invalid session")
	}
}
