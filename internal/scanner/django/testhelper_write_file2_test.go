//ff:func feature=scan type=test control=sequence topic=django
//ff:what writeFile2 테스트 헬퍼
package django

import (
	"os"
	"path/filepath"
	"testing"
)

func writeFile2(t *testing.T, dir, rel, content string) {
	t.Helper()
	p := filepath.Join(dir, rel)
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}
