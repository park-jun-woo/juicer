//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestDetectActix — Cargo.toml에서 actix-web 감지 테스트
package actix

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestDetectActix(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "Cargo.toml", `[package]
name = "my-app"
version = "0.1.0"

[dependencies]
actix-web = "4"
serde = { version = "1", features = ["derive"] }
`)

	fw := scanner.DetectFramework(dir)
	if fw != "actix" {
		t.Errorf("want actix, got %s", fw)
	}
}
