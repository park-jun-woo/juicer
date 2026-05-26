//ff:func feature=sql type=test control=sequence
//ff:what TestRunSkip_WithSessionCov 테스트
package sqls

import (
	"os"
	"testing"
)

func TestRunSkip_WithSessionCov(t *testing.T) {
	dir := t.TempDir()
	sessionDir := dir + "/.juicer"
	os.MkdirAll(sessionDir, 0o755)
	os.WriteFile(sessionDir+"/sql-session.json", []byte(`{"repo_dir":"/tmp","queries_dir":"/tmp","methods":[{"id":"R.M","status":"TODO"}]}`), 0o644)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	RunSkip()
}
