//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what mkFile 테스트 헬퍼
package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func mkFile(t *testing.T, dir, rel, content string) string {
	t.Helper()
	p := filepath.Join(dir, rel)
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}
