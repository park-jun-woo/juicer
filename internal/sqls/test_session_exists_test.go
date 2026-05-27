//ff:func feature=sql type=session control=sequence
//ff:what TestSessionExists 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSessionExists(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	if SessionExists() {
		t.Error("expected false for non-existent session")
	}

	os.MkdirAll(".codist", 0o755)
	os.WriteFile(filepath.Join(".codist", "sql-session.json"), []byte("{}"), 0o644)

	if !SessionExists() {
		t.Error("expected true for existing session")
	}
}
