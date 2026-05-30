//ff:func feature=scan type=test control=sequence topic=django
//ff:what detectDjango — django 의존 감지(flask/fastapi 배제) 분기를 검증
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectDjango_Hit(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte("Django==4.2\n"), 0o644)
	if !detectDjango(dir) {
		t.Fatal("expected true")
	}
}

func TestDetectDjango_ExcludedByFlask(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte("django\nflask\n"), 0o644)
	if detectDjango(dir) {
		t.Fatal("expected false when flask is present")
	}
}

func TestDetectDjango_ExcludedByFastAPI(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte("django\nfastapi\n"), 0o644)
	if detectDjango(dir) {
		t.Fatal("expected false when fastapi is present")
	}
}

func TestDetectDjango_NoFiles(t *testing.T) {
	if detectDjango(t.TempDir()) {
		t.Fatal("expected false when no dependency files")
	}
}
