//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what TestExprName 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestExprName(t *testing.T) {
	tests := []struct {
		name string
		expr ast.Expr
		want string
	}{
		{"ident", &ast.Ident{Name: "foo"}, "foo"},
		{"selector", &ast.SelectorExpr{X: &ast.Ident{Name: "h"}, Sel: &ast.Ident{Name: "Method"}}, "h.Method"},
		{"func lit", &ast.FuncLit{Type: &ast.FuncType{}, Body: &ast.BlockStmt{}}, "(inline)"},
		{"call", &ast.CallExpr{Fun: &ast.Ident{Name: "fn"}}, "fn()"},
		{"selector with non-ident X", &ast.SelectorExpr{X: &ast.CompositeLit{}, Sel: &ast.Ident{Name: "Method"}}, "Method"},
		{"unknown", &ast.CompositeLit{}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := exprName(tt.expr)
			if got != tt.want {
				t.Errorf("exprName() = %q, want %q", got, tt.want)
			}
		})
	}
}
