//ff:func feature=scan type=extract control=sequence
//ff:what TestHandleForm 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestHandleForm(t *testing.T) {
	t.Run("basic form", func(t *testing.T) {
		ep := &scanner.Endpoint{}
		call := &ast.CallExpr{
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: `"field"`},
			},
		}
		handleForm(ep, call)

		if ep.Request == nil {
			t.Fatal("expected request")
		}
		if len(ep.Request.FormFields) != 1 {
			t.Fatalf("expected 1 form field, got %d", len(ep.Request.FormFields))
		}
	})

	t.Run("no args", func(t *testing.T) {
		ep := &scanner.Endpoint{}
		call := &ast.CallExpr{}
		handleForm(ep, call)
		if ep.Request != nil {
			t.Error("expected no request")
		}
	})

	t.Run("non-string arg form", func(t *testing.T) {
		ep := &scanner.Endpoint{}
		call := &ast.CallExpr{
			Args: []ast.Expr{&ast.Ident{Name: "varName"}},
		}
		handleForm(ep, call)
		if ep.Request != nil {
			t.Error("expected no request for non-string arg")
		}
	})

	t.Run("duplicate ignored", func(t *testing.T) {
		ep := &scanner.Endpoint{
			Request: &scanner.Request{
				FormFields: []scanner.Param{{Name: "field"}},
			},
		}
		call := &ast.CallExpr{
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: `"field"`},
			},
		}
		handleForm(ep, call)
		if len(ep.Request.FormFields) != 1 {
			t.Error("duplicate should be ignored")
		}
	})
}
