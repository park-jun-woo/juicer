//ff:func feature=scan type=extract control=sequence
//ff:what TestHandlePathParam 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestHandlePathParam(t *testing.T) {
	t.Run("basic param", func(t *testing.T) {
		ep := &scanner.Endpoint{}
		call := &ast.CallExpr{
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: `"id"`},
			},
		}
		handlePathParam(ep, call)

		if ep.Request == nil {
			t.Fatal("expected request")
		}
		if len(ep.Request.PathParams) != 1 {
			t.Fatalf("expected 1 path param, got %d", len(ep.Request.PathParams))
		}
	})

	t.Run("no args", func(t *testing.T) {
		ep := &scanner.Endpoint{}
		call := &ast.CallExpr{}
		handlePathParam(ep, call)
		if ep.Request != nil {
			t.Error("expected no request")
		}
	})

	t.Run("non-string arg", func(t *testing.T) {
		ep := &scanner.Endpoint{}
		call := &ast.CallExpr{
			Args: []ast.Expr{&ast.Ident{Name: "varName"}},
		}
		handlePathParam(ep, call)
		if ep.Request != nil {
			t.Error("expected no request for non-string arg")
		}
	})

	t.Run("duplicate ignored", func(t *testing.T) {
		ep := &scanner.Endpoint{
			Request: &scanner.Request{
				PathParams: []scanner.Param{{Name: "id"}},
			},
		}
		call := &ast.CallExpr{
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: `"id"`},
			},
		}
		handlePathParam(ep, call)
		if len(ep.Request.PathParams) != 1 {
			t.Error("duplicate should be ignored")
		}
	})
}
