//ff:func feature=sql type=test control=sequence
//ff:what TestRunNext_CreateSessionCov 테스트
package sqls

import (
	"os"
	"testing"
)

func TestRunNext_CreateSessionCov(t *testing.T) {
	dir := t.TempDir()
	repoDir := dir + "/repo"
	queriesDir := dir + "/queries"
	os.MkdirAll(repoDir, 0o755)
	os.MkdirAll(queriesDir, 0o755)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	RunNext(repoDir, queriesDir)
}
