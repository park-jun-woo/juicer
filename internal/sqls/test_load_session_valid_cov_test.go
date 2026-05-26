//ff:func feature=sql type=test control=sequence
//ff:what TestLoadSession_ValidCov 테스트
package sqls

import (
	"os"
	"testing"
)

func TestLoadSession_ValidCov(t *testing.T) {
	dir := t.TempDir()
	sessionDir := dir + "/.juicer"
	os.MkdirAll(sessionDir, 0o755)
	os.WriteFile(sessionDir+"/sql-session.json", []byte(`{"repo_dir":"/tmp","queries_dir":"/tmp","methods":[]}`), 0o644)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	sess, err := LoadSession()
	if err != nil {
		t.Fatal(err)
	}
	if sess == nil {
		t.Fatal("expected session")
	}
}
