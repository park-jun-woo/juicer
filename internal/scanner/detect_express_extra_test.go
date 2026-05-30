//ff:func feature=scan type=test control=sequence
//ff:what detectExpress — express 의존 감지(@nestjs/core 배제) 분기를 검증
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectExpress_Hit(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "package.json"),
		[]byte(`{"dependencies":{"express":"^4.18.0"}}`), 0o644)
	if !detectExpress(dir) {
		t.Fatal("expected true")
	}
}

func TestDetectExpress_ExcludedByNest(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "package.json"),
		[]byte(`{"dependencies":{"express":"^4","@nestjs/core":"^10"}}`), 0o644)
	if detectExpress(dir) {
		t.Fatal("expected false when @nestjs/core present")
	}
}

func TestDetectExpress_NoExpress(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "package.json"), []byte(`{"dependencies":{}}`), 0o644)
	if detectExpress(dir) {
		t.Fatal("expected false without express")
	}
}

func TestDetectExpress_NoFile(t *testing.T) {
	if detectExpress(t.TempDir()) {
		t.Fatal("expected false when package.json missing")
	}
}
