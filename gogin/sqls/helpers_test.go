//ff:func feature=sql type=session control=sequence
//ff:what 테스트 헬퍼 — 세션 디렉토리 설정
package sqls

import (
	"os"
	"testing"
)

func setupSessionDir(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	t.Cleanup(func() { os.Chdir(oldWd) })
	return dir
}
