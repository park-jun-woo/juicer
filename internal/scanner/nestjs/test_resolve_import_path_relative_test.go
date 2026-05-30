//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveImportPath_Relative 테스트
package nestjs

import (
	"path/filepath"
	"testing"
)

func TestResolveImportPath_Relative(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "dto/create.dto.ts", "x")
	got := resolveImportPath(dir, "./dto/create.dto")
	if got != filepath.Join(dir, "dto/create.dto.ts") {
		t.Fatalf("got %q", got)
	}
}
