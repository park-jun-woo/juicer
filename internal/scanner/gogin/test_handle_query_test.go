//ff:func feature=scan type=extract control=sequence
//ff:what TestHandleQuery 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleQuery(t *testing.T) {
	t.Run("basic query", func(t *testing.T) {
		ep := &scanner.Endpoint{}
		call := &ast.CallExpr{
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: `"name"`},
			},
		}
		handleQuery(ep, call, "Query")

		if ep.Request == nil {
			t.Fatal("expected request to be set")
		}
		if len(ep.Request.Query) != 1 {
			t.Fatalf("expected 1 query, got %d", len(ep.Request.Query))
		}
		if ep.Request.Query[0].Name != "name" {
			t.Errorf("expected name 'name', got %q", ep.Request.Query[0].Name)
		}
	})

	t.Run("DefaultQuery with default", func(t *testing.T) {
		ep := &scanner.Endpoint{}
		call := &ast.CallExpr{
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: `"page"`},
				&ast.BasicLit{Kind: token.STRING, Value: `"1"`},
			},
		}
		handleQuery(ep, call, "DefaultQuery")

		if ep.Request.Query[0].Default != "1" {
			t.Errorf("expected default '1', got %q", ep.Request.Query[0].Default)
		}
	})

	t.Run("no args", func(t *testing.T) {
		ep := &scanner.Endpoint{}
		call := &ast.CallExpr{}
		handleQuery(ep, call, "Query")
		if ep.Request != nil {
			t.Error("expected no request for no args")
		}
	})

	t.Run("non-string arg", func(t *testing.T) {
		ep := &scanner.Endpoint{}
		call := &ast.CallExpr{
			Args: []ast.Expr{&ast.Ident{Name: "varName"}},
		}
		handleQuery(ep, call, "Query")
		if ep.Request != nil {
			t.Error("expected no request for non-string arg")
		}
	})

	t.Run("duplicate ignored", func(t *testing.T) {
		ep := &scanner.Endpoint{
			Request: &scanner.Request{
				Query: []scanner.Param{{Name: "name"}},
			},
		}
		call := &ast.CallExpr{
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: `"name"`},
			},
		}
		handleQuery(ep, call, "Query")
		if len(ep.Request.Query) != 1 {
			t.Error("duplicate query should be ignored")
		}
	})
}
