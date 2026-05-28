//ff:func feature=scan type=test control=sequence
//ff:what TestHandleResponse_SendStatus 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleResponse_SendStatus(t *testing.T) {
	ep := scanner.Endpoint{}
	call := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "c"},
			Sel: &ast.Ident{Name: "SendStatus"},
		},
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.INT, Value: "204"},
		},
	}

	handleResponse(&ep, call, "status", nil, "handler")

	if len(ep.Responses) != 1 {
		t.Fatalf("expected 1 response, got %d", len(ep.Responses))
	}
	if ep.Responses[0].Status != "204" {
		t.Fatalf("expected status 204, got %s", ep.Responses[0].Status)
	}
}
