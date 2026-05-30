//ff:func feature=scan type=test control=sequence
//ff:what TestHandleChainedResponse_StatusJSON 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"testing"
)

func TestHandleChainedResponse_StatusJSON(t *testing.T) {
	ep := scanner.Endpoint{}

	statusCall := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "c"},
			Sel: &ast.Ident{Name: "Status"},
		},
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.INT, Value: "201"},
		},
	}

	outerCall := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   statusCall,
			Sel: &ast.Ident{Name: "JSON"},
		},
		Args: []ast.Expr{
			&ast.Ident{Name: "result"},
		},
	}

	handleChainedResponse(&ep, statusCall, outerCall, "JSON", nil, "handler")

	if len(ep.Responses) != 1 {
		t.Fatalf("expected 1 response, got %d", len(ep.Responses))
	}
	if ep.Responses[0].Status != "201" {
		t.Fatalf("expected status 201, got %s", ep.Responses[0].Status)
	}
	if ep.Responses[0].Kind != "json" {
		t.Fatalf("expected kind json, got %s", ep.Responses[0].Kind)
	}
	if ep.Responses[0].Body != "result" {
		t.Fatalf("expected body 'result', got %s", ep.Responses[0].Body)
	}
}
