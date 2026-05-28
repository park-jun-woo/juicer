//ff:func feature=scan type=test control=sequence topic=express
//ff:what package.json 없을 때 Express 감지 실패 확인
package express

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestDetectExpress_NoPackageJson(t *testing.T) {
	dir := t.TempDir()
	fw := scanner.DetectFramework(dir)
	if fw == "express" {
		t.Errorf("should not detect express without package.json")
	}
}
