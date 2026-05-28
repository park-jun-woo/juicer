//ff:func feature=scan type=test control=sequence topic=express
//ff:what Express 프레임워크 감지 테스트: package.json에 express 의존 확인
package express

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestDetectExpress_WithExpress(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "package.json", `{
  "dependencies": {
    "express": "^4.18.0"
  }
}`)
	fw := scanner.DetectFramework(dir)
	if fw != "express" {
		t.Errorf("want express, got %s", fw)
	}
}
