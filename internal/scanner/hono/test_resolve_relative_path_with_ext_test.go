//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestResolveRelativePath_WithExt 테스트
package hono

import (
	"path/filepath"
	"testing"
)

func TestResolveRelativePath_WithExt(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "mod.ts", "x\n")
	if got := resolveRelativePath(dir, "./mod"); got != filepath.Join(dir, "mod.ts") {
		t.Fatalf("got %q", got)
	}
}
