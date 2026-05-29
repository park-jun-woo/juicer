//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestDetectActix_NotFound — actix가 아닌 의존성에서 미감지 확인
package actix

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestDetectActix_NotFound(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "Cargo.toml", `[package]
name = "my-app"

[dependencies]
rocket = "0.5"
`)

	fw := scanner.DetectFramework(dir)
	if fw == "actix" {
		t.Errorf("should not detect actix")
	}
}
