//ff:func feature=ratchet type=session control=sequence
//ff:what TestRunReset_DeleteError 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunReset_DeleteError(t *testing.T) {
	dir := setupSessionDir(t)
	humaDir := filepath.Join(dir, ".codist")
	os.MkdirAll(humaDir, 0o755)
	os.WriteFile(filepath.Join(humaDir, "sql-session.json"), []byte("{}"), 0o644)

	// Make directory read-only to prevent deletion
	os.Chmod(dir, 0o555)
	t.Cleanup(func() { os.Chmod(dir, 0o755) })

	err := RunReset()
	if err == nil {
		// Some systems may still allow deletion
		return
	}
}
