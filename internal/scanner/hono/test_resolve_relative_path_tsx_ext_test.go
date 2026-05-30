//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestResolveRelativePath_TsxExt 테스트
package hono

import (
	"path/filepath"
	"testing"
)

func TestResolveRelativePath_TsxExt(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "comp.tsx", "x\n")
	if got := resolveRelativePath(dir, "./comp"); got != filepath.Join(dir, "comp.tsx") {
		t.Fatalf("got %q", got)
	}
}
