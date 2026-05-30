//ff:func feature=scan type=test control=sequence
//ff:what TestHandleForm_Basic 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleForm_Basic(t *testing.T) {
	ep := scanner.Endpoint{}
	call := &ast.CallExpr{
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.STRING, Value: `"name"`},
		},
	}

	handleForm(&ep, call)

	if ep.Request == nil {
		t.Fatal("expected non-nil Request")
	}
	if len(ep.Request.FormFields) != 1 || ep.Request.FormFields[0].Name != "name" {
		t.Fatalf("expected form field 'name', got %v", ep.Request.FormFields)
	}
}

func TestHandleForm_NoArgs(t *testing.T) {
	ep := scanner.Endpoint{}
	handleForm(&ep, &ast.CallExpr{})
	if ep.Request != nil {
		t.Fatalf("expected no request for no args, got %v", ep.Request)
	}
}

func TestHandleForm_EmptyName(t *testing.T) {
	ep := scanner.Endpoint{}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "v"}}}
	handleForm(&ep, call)
	if ep.Request != nil {
		t.Fatalf("expected no request for empty name, got %v", ep.Request)
	}
}

func TestHandleForm_Duplicate(t *testing.T) {
	ep := scanner.Endpoint{}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"name"`}}}
	handleForm(&ep, call)
	handleForm(&ep, call)
	if len(ep.Request.FormFields) != 1 {
		t.Fatalf("expected 1 form field (dedup), got %d", len(ep.Request.FormFields))
	}
}
