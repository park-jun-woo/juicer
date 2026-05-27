//ff:func feature=sql type=test control=sequence
//ff:what TestSessionExists_WithFileCov 테스트
package sqls

import (
	"os"
	"testing"
)

func TestSessionExists_WithFileCov(t *testing.T) {
	dir := t.TempDir()
	sessionDir := dir + "/.codist"
	os.MkdirAll(sessionDir, 0o755)
	os.WriteFile(sessionDir+"/sql-session.json", []byte("{}"), 0o644)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	if !SessionExists() {
		t.Fatal("expected true")
	}
}
