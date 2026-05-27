//ff:func feature=ratchet type=session control=sequence
//ff:what TestRunSkip_LoadError 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunSkip_LoadError(t *testing.T) {
	dir := setupSessionDir(t)
	os.MkdirAll(filepath.Join(dir, ".codist"), 0o755)
	os.WriteFile(filepath.Join(dir, ".codist", "sql-session.json"), []byte("bad"), 0o644)

	err := RunSkip()
	if err == nil {
		t.Error("expected error for invalid session")
	}
}
