//ff:func feature=sql type=test control=sequence
//ff:what TestRunSkip_NoSessionCov 테스트
package sqls

import (
	"os"
	"testing"
)

func TestRunSkip_NoSessionCov(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	RunSkip()
}
