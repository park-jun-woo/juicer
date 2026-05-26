//ff:func feature=sql type=test control=sequence
//ff:what TestRunNext_WithSessionCov 테스트
package sqls

import (
	"os"
	"testing"
)

func TestRunNext_WithSessionCov(t *testing.T) {
	dir := t.TempDir()
	repoDir := dir + "/repo"
	queriesDir := dir + "/queries"
	os.MkdirAll(repoDir, 0o755)
	os.MkdirAll(queriesDir, 0o755)
	sessionDir := dir + "/.juicer"
	os.MkdirAll(sessionDir, 0o755)
	os.WriteFile(sessionDir+"/sql-session.json", []byte(`{"repo_dir":"`+repoDir+`","queries_dir":"`+queriesDir+`","methods":[]}`), 0o644)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	err := RunNext("", "")
	if err != nil {
		t.Fatal(err)
	}
}
