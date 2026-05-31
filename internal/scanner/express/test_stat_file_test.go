//ff:func feature=scan type=test control=sequence topic=express
//ff:what statFile 파일 존재 여부 테스트
package express

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStatFile(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "a.ts")
	if err := os.WriteFile(f, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if !statFile(f) {
		t.Error("existing file should stat true")
	}
	if statFile(filepath.Join(dir, "nope.ts")) {
		t.Error("missing file should be false")
	}
}
