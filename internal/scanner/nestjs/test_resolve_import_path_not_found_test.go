//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveImportPath_NotFound 테스트
package nestjs

import "testing"

func TestResolveImportPath_NotFound(t *testing.T) {
	dir := t.TempDir()
	result := resolveImportPath(dir, "./nonexistent")
	if result != "" {
		t.Fatal("expected empty for missing file")
	}
}
