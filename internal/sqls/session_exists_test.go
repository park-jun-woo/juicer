//ff:func feature=sql type=test control=sequence
//ff:what TestSessionExists_NoFile 테스트
package sqls

import (
	"os"
	"testing"
)

func TestSessionExists_NoFile(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	if SessionExists() {
		t.Fatal("expected false in empty dir")
	}
}
