//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildRequest_WithFiles 테스트
package nestjs

import (
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestBuildRequest_WithFiles(t *testing.T) {
	ep := endpointInfo{files: []scanner.Param{{Name: "avatar", Type: "file"}}}
	req := buildRequest(ep)
	if req == nil || len(req.Files) != 1 {
		t.Fatal("expected file params")
	}
}
