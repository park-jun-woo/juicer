//ff:func feature=scan type=extract control=sequence
//ff:what TestHandleBind 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestHandleBind(t *testing.T) {
	t.Run("basic bind", func(t *testing.T) {
		ep := &Endpoint{}
		call := &ast.CallExpr{
			Args: []ast.Expr{
				&ast.UnaryExpr{
					Op: token.AND,
					X:  &ast.Ident{Name: "req"},
				},
			},
		}
		info := &types.Info{
			Types: make(map[ast.Expr]types.TypeAndValue),
			Uses:  make(map[*ast.Ident]types.Object),
		}
		handleBind(ep, call, "ShouldBindJSON", info)

		if ep.Request == nil || ep.Request.Body == nil {
			t.Fatal("expected body to be set")
		}
		if ep.Request.Body.VarName != "req" {
			t.Errorf("expected VarName 'req', got %q", ep.Request.Body.VarName)
		}
		if ep.Request.Body.Method != "ShouldBindJSON" {
			t.Errorf("expected Method 'ShouldBindJSON', got %q", ep.Request.Body.Method)
		}
	})

	t.Run("no args", func(t *testing.T) {
		ep := &Endpoint{}
		call := &ast.CallExpr{}
		info := &types.Info{
			Types: make(map[ast.Expr]types.TypeAndValue),
			Uses:  make(map[*ast.Ident]types.Object),
		}
		handleBind(ep, call, "Bind", info)

		if ep.Request == nil || ep.Request.Body == nil {
			t.Fatal("expected body to be set")
		}
		if ep.Request.Body.VarName != "(unknown)" {
			t.Errorf("expected VarName '(unknown)', got %q", ep.Request.Body.VarName)
		}
	})

	t.Run("second bind ignored", func(t *testing.T) {
		ep := &Endpoint{
			Request: &Request{
				Body: &Body{VarName: "first"},
			},
		}
		call := &ast.CallExpr{
			Args: []ast.Expr{&ast.Ident{Name: "second"}},
		}
		info := &types.Info{
			Types: make(map[ast.Expr]types.TypeAndValue),
			Uses:  make(map[*ast.Ident]types.Object),
		}
		handleBind(ep, call, "Bind", info)

		if ep.Request.Body.VarName != "first" {
			t.Error("second bind should be ignored")
		}
	})
}
