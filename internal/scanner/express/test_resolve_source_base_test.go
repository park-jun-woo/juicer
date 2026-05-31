//ff:func feature=scan type=test topic=express control=sequence
//ff:what resolveSourceBase 확장자 없는 base 후보/직접 stat/없음 테스트
package express

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveSourceBase(t *testing.T) {
	dir := t.TempDir()
	// base without extension -> resolves to .ts candidate
	base := filepath.Join(dir, "mod")
	if err := os.WriteFile(base+".ts", []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if got := resolveSourceBase(base); got != base+".ts" {
		t.Errorf("candidate: got %q", got)
	}
	// base already with extension and existing -> returned as-is
	withExt := filepath.Join(dir, "a.js")
	if err := os.WriteFile(withExt, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if got := resolveSourceBase(withExt); got != withExt {
		t.Errorf("direct: got %q", got)
	}
	// nothing
	if got := resolveSourceBase(filepath.Join(dir, "ghost")); got != "" {
		t.Errorf("none: got %q", got)
	}
}
