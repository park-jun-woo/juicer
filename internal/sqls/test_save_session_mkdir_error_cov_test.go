//ff:func feature=sql type=test control=sequence
//ff:what TestSaveSession_MkdirErrorCov 테스트
package sqls

import (
	"os"
	"testing"
)

func TestSaveSession_MkdirErrorCov(t *testing.T) {
	dir := t.TempDir()
	block := dir + "/block"
	os.WriteFile(block, []byte("x"), 0o644)
	origDir, _ := os.Getwd()
	os.Chdir(block)
	defer os.Chdir(origDir)
	sess := &Session{}
	SaveSession(sess)
}
