//ff:func feature=scan type=test control=sequence topic=express
//ff:what NestJS 프로젝트에서 Express로 감지되지 않는지 확인
package express

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestDetectExpress_WithNestJS(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "package.json", `{
  "dependencies": {
    "express": "^4.18.0",
    "@nestjs/core": "^10.0.0"
  }
}`)
	fw := scanner.DetectFramework(dir)
	if fw == "express" {
		t.Errorf("should detect nestjs, not express; got %s", fw)
	}
}
