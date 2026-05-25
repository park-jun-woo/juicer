//ff:func feature=ratchet type=session control=sequence
//ff:what TestRunStatus_LoadError 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunStatus_LoadError(t *testing.T) {
	dir := setupSessionDir(t)
	os.MkdirAll(filepath.Join(dir, ".huma"), 0o755)
	os.WriteFile(filepath.Join(dir, ".huma", "sql-session.json"), []byte("bad"), 0o644)

	err := RunStatus()
	if err == nil {
		t.Error("expected error for invalid session")
	}
}
