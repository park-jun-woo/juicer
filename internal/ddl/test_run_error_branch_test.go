//ff:func feature=ddl type=test control=sequence
//ff:what TestRun_ErrorBranch 테스트
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRun_ErrorBranch(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "nonexistent")
	// Create an unreadable .up.sql file to trigger Parse error
	os.MkdirAll(dir, 0o755)
	f := filepath.Join(dir, "001.up.sql")
	os.WriteFile(f, []byte("data"), 0o644)
	os.Chmod(f, 0o000)
	defer os.Chmod(f, 0o644)
	_, err := Run(dir)
	if err == nil {
		t.Fatal("expected error")
	}
}
