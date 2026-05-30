//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveExtension_AddsExtension 테스트
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
