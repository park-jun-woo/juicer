//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestCaptureResponse_Round5 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestCaptureResponse_Round5(t *testing.T) {
	fi := aFi(t, `async fn h() -> impl Responder { HttpResponse::Ok().json(user) }`)
	block := aFirst(t, fi.root, "block")
	ctx := &responseCtx{
		block:     block,
		src:       fi.src,
		sIdx:      structIndex{},
		cache:     map[string][]scanner.Field{},
		seen:      map[string]bool{},
		responses: nil,
	}
	walkNodes(block, func(n *sitter.Node) {
		captureResponse(n, ctx)
	})
}
