//ff:func feature=scan type=test control=sequence
//ff:what detectFlask — flask 의존 감지(fastapi 배제) 분기를 검증
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectFlask_Hit(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte("Flask==3.0\n"), 0o644)
	if !detectFlask(dir) {
		t.Fatal("expected true")
	}
}

func TestDetectFlask_ExcludedByFastAPI(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte("flask\nfastapi\n"), 0o644)
	if detectFlask(dir) {
		t.Fatal("expected false when fastapi present")
	}
}

func TestDetectFlask_NoFiles(t *testing.T) {
	if detectFlask(t.TempDir()) {
		t.Fatal("expected false when no dependency files")
	}
}
