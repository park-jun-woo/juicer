//ff:func feature=scan type=test control=sequence
//ff:what TestHandleResponse_SourceAndUnknownStatus 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"testing"
)

func TestHandleResponse_SourceAndUnknownStatus(t *testing.T) {
	ep := scanner.Endpoint{}

	call := &ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "c"}, Sel: &ast.Ident{Name: "SendStatus"}}}
	handleResponse(&ep, call, "status", nil, "respond")
	if ep.Responses[0].Status != "(unknown)" || ep.Responses[0].Source != "respond" {
		t.Fatalf("source/unknown: %+v", ep.Responses[0])
	}
}
