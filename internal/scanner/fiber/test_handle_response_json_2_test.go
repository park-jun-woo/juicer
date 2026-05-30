//ff:func feature=scan type=test control=sequence
//ff:what TestHandleResponse_JSON 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"testing"
)

func TestHandleResponse_JSON(t *testing.T) {
	ep := scanner.Endpoint{}
	call := &ast.CallExpr{
		Args: []ast.Expr{
			&ast.Ident{Name: "result"},
		},
	}

	handleResponse(&ep, call, "json", nil, "handler")

	if len(ep.Responses) != 1 {
		t.Fatalf("expected 1 response, got %d", len(ep.Responses))
	}
	if ep.Responses[0].Status != "200" {
		t.Fatalf("expected status 200, got %s", ep.Responses[0].Status)
	}
	if ep.Responses[0].Kind != "json" {
		t.Fatalf("expected kind json, got %s", ep.Responses[0].Kind)
	}
}
