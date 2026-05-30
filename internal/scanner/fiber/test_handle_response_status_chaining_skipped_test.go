//ff:func feature=scan type=test control=sequence
//ff:what TestHandleResponse_StatusChainingSkipped 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"testing"
)

func TestHandleResponse_StatusChainingSkipped(t *testing.T) {
	ep := scanner.Endpoint{}

	call := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "c"}, Sel: &ast.Ident{Name: "Status"}},
		Args: []ast.Expr{&ast.BasicLit{Kind: token.INT, Value: "200"}},
	}
	handleResponse(&ep, call, "status", nil, "handler")
	if len(ep.Responses) != 0 {
		t.Fatalf("c.Status() should not record a response, got %v", ep.Responses)
	}
}
