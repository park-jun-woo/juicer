//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveExtension_NotFound 테스트
package express

import (
	"path/filepath"
	"testing"
)

func TestResolveExtension_NotFound(t *testing.T) {
	dir := t.TempDir()
	if got := resolveExtension(filepath.Join(dir, "nope")); got != "" {
		t.Fatalf("got %q", got)
	}
}
