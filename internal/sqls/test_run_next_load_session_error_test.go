//ff:func feature=ratchet type=session control=sequence
//ff:what TestRunNext_LoadSessionError 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunNext_LoadSessionError(t *testing.T) {
	dir := setupSessionDir(t)
	// Create session file with invalid JSON
	os.MkdirAll(filepath.Join(dir, ".juicer"), 0o755)
	os.WriteFile(filepath.Join(dir, ".juicer", "sql-session.json"), []byte("invalid json"), 0o644)

	err := RunNext("", "")
	if err == nil {
		t.Error("expected error for invalid session")
	}
}
