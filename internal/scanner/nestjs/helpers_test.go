//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what 테스트 헬퍼 함수
package nestjs

import (
	"os"
	"path/filepath"
	"testing"
)

// writeFile creates a file with the given content in the temp dir.
func writeFile(t *testing.T, dir, rel, content string) {
	t.Helper()
	p := filepath.Join(dir, rel)
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}
