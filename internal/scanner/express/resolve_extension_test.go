//ff:func feature=scan type=test control=sequence topic=express
//ff:what resolveExtension: 확장자 추가 해석 / 미존재 빈문자열
package express

import (
	"path/filepath"
	"testing"
)

func TestResolveExtension_AddsExtension(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes.ts", "x")
	base := filepath.Join(dir, "routes")
	got := resolveExtension(base)
	if got != filepath.Join(dir, "routes.ts") {
		t.Fatalf("got %q", got)
	}
}

func TestResolveExtension_NotFound(t *testing.T) {
	dir := t.TempDir()
	if got := resolveExtension(filepath.Join(dir, "nope")); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestResolveExtension_ExistingWithExt(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "a.ts", "x")
	base := filepath.Join(dir, "a.ts")
	if got := resolveExtension(base); got != base {
		t.Fatalf("got %q", got)
	}
}
