//ff:func feature=sql type=parse control=sequence
//ff:what TestRunReset 테스트
package sqls

import (
	"os"
	"testing"
)

func TestRunReset_Cov(t *testing.T) {
	dir := t.TempDir()
	sessionDir := dir + "/.codist"
	os.MkdirAll(sessionDir, 0o755)
	os.WriteFile(sessionDir+"/sql-session.json", []byte(`{}`), 0o644)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	err := RunReset()
	if err != nil {
		t.Fatal(err)
	}
}

