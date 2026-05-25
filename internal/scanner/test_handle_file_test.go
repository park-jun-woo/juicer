//ff:func feature=scan type=extract control=sequence
//ff:what TestHandleFile 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestHandleFile(t *testing.T) {
	t.Run("basic file", func(t *testing.T) {
		ep := &Endpoint{}
		call := &ast.CallExpr{
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: `"avatar"`},
			},
		}
		handleFile(ep, call)

		if ep.Request == nil {
			t.Fatal("expected request")
		}
		if len(ep.Request.Files) != 1 {
			t.Fatalf("expected 1 file, got %d", len(ep.Request.Files))
		}
	})

	t.Run("no args", func(t *testing.T) {
		ep := &Endpoint{}
		call := &ast.CallExpr{}
		handleFile(ep, call)
		if ep.Request != nil {
			t.Error("expected no request")
		}
	})

	t.Run("non-string arg file", func(t *testing.T) {
		ep := &Endpoint{}
		call := &ast.CallExpr{
			Args: []ast.Expr{&ast.Ident{Name: "varName"}},
		}
		handleFile(ep, call)
		if ep.Request != nil {
			t.Error("expected no request for non-string arg")
		}
	})

	t.Run("duplicate ignored", func(t *testing.T) {
		ep := &Endpoint{
			Request: &Request{
				Files: []Param{{Name: "avatar"}},
			},
		}
		call := &ast.CallExpr{
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: `"avatar"`},
			},
		}
		handleFile(ep, call)
		if len(ep.Request.Files) != 1 {
			t.Error("duplicate should be ignored")
		}
	})
}
