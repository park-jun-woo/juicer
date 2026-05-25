//ff:func feature=sql type=parse control=sequence
//ff:what TestSaveSession_Basic 테스트
package sqls

import (
	"os"
	"testing"
)

func TestSaveSession_Basic(t *testing.T) {
	dir, err := os.MkdirTemp("", "sqls-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	origDir, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(origDir)

	sess := &Session{RepoDir: dir, Methods: []MethodStatus{}}
	err = SaveSession(sess)
	if err != nil {
		t.Fatal(err)
	}
}
