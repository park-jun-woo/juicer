//ff:func feature=ddl type=parse control=sequence
//ff:what TestRun_ParseError 테스트
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRun_ParseError(t *testing.T) {
	dir := t.TempDir()
	// Create a .up.sql file that is unreadable
	path := filepath.Join(dir, "001_init.up.sql")
	if err := os.WriteFile(path, []byte("SELECT 1"), 0o644); err != nil {
		t.Fatal(err)
	}
	// Remove read permissions
	if err := os.Chmod(path, 0o000); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		os.Chmod(path, 0o644) // restore for cleanup
	})

	_, err := Run(dir)
	if err == nil {
		t.Fatal("Run() expected error for unreadable file, got nil")
	}
}
