//ff:func feature=sql type=test control=sequence
//ff:what TestLoadSession_InvalidJSONCov 테스트
package sqls

import (
	"os"
	"testing"
)

func TestLoadSession_InvalidJSONCov(t *testing.T) {
	dir := t.TempDir()
	sessionDir := dir + "/.juicer"
	os.MkdirAll(sessionDir, 0o755)
	os.WriteFile(sessionDir+"/sql-session.json", []byte("{invalid"), 0o644)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	_, err := LoadSession()
	if err == nil {
		t.Fatal("expected error")
	}
}

