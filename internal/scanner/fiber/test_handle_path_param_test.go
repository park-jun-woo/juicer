//ff:func feature=scan type=test control=sequence
//ff:what TestHandlePathParam_Basic 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandlePathParam_Basic(t *testing.T) {
	ep := scanner.Endpoint{}
	call := &ast.CallExpr{
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.STRING, Value: `"id"`},
		},
	}

	handlePathParam(&ep, call)

	if ep.Request == nil {
		t.Fatal("expected non-nil Request")
	}
	if len(ep.Request.PathParams) != 1 || ep.Request.PathParams[0].Name != "id" {
		t.Fatalf("expected path param 'id', got %v", ep.Request.PathParams)
	}
}

func TestHandlePathParam_NoArgs(t *testing.T) {
	ep := scanner.Endpoint{}
	handlePathParam(&ep, &ast.CallExpr{})
	if ep.Request != nil {
		t.Fatalf("expected no request, got %v", ep.Request)
	}
}

func TestHandlePathParam_EmptyName(t *testing.T) {
	ep := scanner.Endpoint{}
	handlePathParam(&ep, &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "v"}}})
	if ep.Request != nil {
		t.Fatalf("expected no request, got %v", ep.Request)
	}
}

func TestHandlePathParam_Duplicate(t *testing.T) {
	ep := scanner.Endpoint{}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"id"`}}}
	handlePathParam(&ep, call)
	handlePathParam(&ep, call)
	if len(ep.Request.PathParams) != 1 {
		t.Fatalf("expected 1 path param (dedup), got %d", len(ep.Request.PathParams))
	}
}
