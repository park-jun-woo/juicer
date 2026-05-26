//ff:func feature=scan type=test control=sequence
//ff:what TestDetectGoGin_NoMatch 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectGoGin_NoMatch(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte("require github.com/gorilla/mux v1.8.0"), 0o644)
	if detectGoGin(dir) {
		t.Fatal("expected false when gin not present")
	}
}
