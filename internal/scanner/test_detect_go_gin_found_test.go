//ff:func feature=scan type=test control=sequence
//ff:what TestDetectGoGin_Found go.mod gin 감지 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectGoGin_Found(t *testing.T) {
	dir := t.TempDir()
	goMod := "module test\nrequire github.com/gin-gonic/gin v1.9.1\n"
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte(goMod), 0o644)
	if !detectGoGin(dir) {
		t.Fatal("expected true")
	}
}
