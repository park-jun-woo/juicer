//ff:func feature=ddl type=parse control=sequence
//ff:what TestParse_ReadError 테스트
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParse_ReadError(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "001_init.up.sql")
	os.WriteFile(path, []byte("SELECT 1"), 0o644)
	os.Chmod(path, 0o000)
	t.Cleanup(func() { os.Chmod(path, 0o644) })

	_, err := Parse(dir)
	if err == nil {
		t.Fatal("Parse() expected error for unreadable file, got nil")
	}
}
