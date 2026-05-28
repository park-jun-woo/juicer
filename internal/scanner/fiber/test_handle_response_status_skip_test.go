//ff:func feature=scan type=test control=sequence
//ff:what TestHandleResponse_StatusSkip 테스트 — c.Status() 단독은 응답으로 기록하지 않는다
package fiber

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleResponse_StatusSkip(t *testing.T) {
	ep := scanner.Endpoint{}
	call := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "c"},
			Sel: &ast.Ident{Name: "Status"},
		},
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.INT, Value: "201"},
		},
	}

	handleResponse(&ep, call, "status", nil, "handler")

	if len(ep.Responses) != 0 {
		t.Fatalf("expected 0 responses (Status is chaining, not terminal), got %d", len(ep.Responses))
	}
}
