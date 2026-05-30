//ff:func feature=scan type=test control=sequence
//ff:what scanWriteFile 테스트 헬퍼
package gogin

import (
	"os"
	"testing"
)

func scanWriteFile(t *testing.T, dir, rel, content string) {
	t.Helper()
	p := scanJoin(dir, rel)
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}
