//ff:func feature=sql type=session control=sequence
//ff:what TestSaveSession_WriteError 테스트
package sqls

import (
	"os"
	"testing"
)

func TestSaveSession_WriteError(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	os.MkdirAll(".juicer", 0o755)
	os.Chmod(".juicer", 0o555)
	defer os.Chmod(".juicer", 0o755)

	sess := &Session{RepoDir: "repo"}
	err := SaveSession(sess)
	if err == nil {
		t.Error("expected error for read-only directory")
	}
}
