//ff:func feature=sql type=test control=sequence
//ff:what TestCreateSession_ValidCov 테스트
package sqls

import (
	"os"
	"testing"
)

func TestCreateSession_ValidCov(t *testing.T) {
	dir := t.TempDir()
	repoDir := dir + "/repo"
	queriesDir := dir + "/queries"
	os.MkdirAll(repoDir, 0o755)
	os.MkdirAll(queriesDir, 0o755)
	sessionDir := dir + "/.juicer"
	os.MkdirAll(sessionDir, 0o755)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	err := createSession(repoDir, queriesDir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
