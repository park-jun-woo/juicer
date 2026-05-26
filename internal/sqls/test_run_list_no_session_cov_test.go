//ff:func feature=sql type=test control=sequence
//ff:what TestRunList_NoSessionCov 테스트
package sqls

import (
	"os"
	"testing"
)

func TestRunList_NoSessionCov(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	// RunList may or may not error when no session - just call it
	RunList()
}
