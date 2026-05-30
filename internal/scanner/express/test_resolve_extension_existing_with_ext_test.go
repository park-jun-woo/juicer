//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveExtension_ExistingWithExt 테스트
package express

import (
	"path/filepath"
	"testing"
)

func TestResolveExtension_ExistingWithExt(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "a.ts", "x")
	base := filepath.Join(dir, "a.ts")
	if got := resolveExtension(base); got != base {
		t.Fatalf("got %q", got)
	}
}
