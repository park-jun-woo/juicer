//ff:func feature=ratchet type=session control=sequence
//ff:what TestRunSqlcGenerate_InTempDir 테스트
package sqls

import (
	"os"
	"testing"
)

func TestRunSqlcGenerate_InTempDir(t *testing.T) {
	// Run in temp dir where sqlc will fail (no config)
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	passed, _ := runSqlcGenerate()
	if passed {
		t.Error("expected sqlc to fail without config")
	}
}
