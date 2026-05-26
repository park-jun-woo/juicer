//ff:func feature=sql type=test control=sequence
//ff:what TestDeleteSession_ExistingCov 테스트
package sqls

import (
	"os"
	"testing"
)

func TestDeleteSession_ExistingCov(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	sessionDir := dir + "/.juicer"
	os.MkdirAll(sessionDir, 0o755)
	os.WriteFile(sessionDir+"/sql-session.json", []byte("{}"), 0o644)
	err := DeleteSession()
	if err != nil {
		t.Fatal(err)
	}
}
