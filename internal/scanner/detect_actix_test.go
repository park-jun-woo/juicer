//ff:func feature=scan type=test control=sequence
//ff:what detectActix — Cargo.toml의 actix-web 의존 감지 분기를 검증
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectActix_Hit(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "Cargo.toml"), []byte("[dependencies]\nactix-web = \"4\"\n"), 0o644)
	if !detectActix(dir) {
		t.Fatal("expected true")
	}
}

func TestDetectActix_Miss(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "Cargo.toml"), []byte("[dependencies]\nrocket = \"0.5\"\n"), 0o644)
	if detectActix(dir) {
		t.Fatal("expected false when actix-web absent")
	}
}

func TestDetectActix_NoFile(t *testing.T) {
	dir := t.TempDir()
	if detectActix(dir) {
		t.Fatal("expected false when Cargo.toml missing")
	}
}
