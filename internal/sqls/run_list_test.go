//ff:func feature=sql type=test control=sequence
//ff:what TestRunList_WithSessionCov 테스트
package sqls

import (
	"os"
	"testing"
)

func TestRunList_WithSessionCov(t *testing.T) {
	dir := t.TempDir()
	sessionDir := dir + "/.codist"
	os.MkdirAll(sessionDir, 0o755)
	os.WriteFile(sessionDir+"/sql-session.json", []byte(`{"repo_dir":"/tmp","queries_dir":"/tmp","methods":[{"id":"Repo.M","status":"TODO"}]}`), 0o644)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	err := RunList()
	if err != nil {
		t.Fatal(err)
	}
}

