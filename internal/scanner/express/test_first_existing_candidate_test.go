//ff:func feature=scan type=test topic=express control=sequence
//ff:what firstExistingCandidate 확장자/인덱스 후보 해석 테스트
package express

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFirstExistingCandidate(t *testing.T) {
	dir := t.TempDir()
	// base + .ts candidate
	base := filepath.Join(dir, "routes")
	if err := os.WriteFile(base+".ts", []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if got := firstExistingCandidate(base); got != base+".ts" {
		t.Errorf("ext candidate: got %q", got)
	}
	// directory index candidate
	d2 := filepath.Join(dir, "pkg")
	if err := os.Mkdir(d2, 0o755); err != nil {
		t.Fatal(err)
	}
	idx := filepath.Join(d2, "index.js")
	if err := os.WriteFile(idx, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if got := firstExistingCandidate(d2); got != idx {
		t.Errorf("index candidate: got %q", got)
	}
	// none
	if got := firstExistingCandidate(filepath.Join(dir, "ghost")); got != "" {
		t.Errorf("none: got %q", got)
	}
}
