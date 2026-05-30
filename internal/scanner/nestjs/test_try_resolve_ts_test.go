//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestTryResolveTS 테스트
package nestjs

import (
	"path/filepath"
	"testing"
)

func TestTryResolveTS(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "mod.ts", "x")
	if got := tryResolveTS(filepath.Join(dir, "mod")); got != filepath.Join(dir, "mod.ts") {
		t.Fatalf("got %q", got)
	}
	writeFile(t, dir, "pkg/index.ts", "x")
	if got := tryResolveTS(filepath.Join(dir, "pkg")); got != filepath.Join(dir, "pkg/index.ts") {
		t.Fatalf("index: %q", got)
	}
	if got := tryResolveTS(filepath.Join(dir, "missing")); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
