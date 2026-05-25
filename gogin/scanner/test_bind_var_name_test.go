//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what TestBindVarName 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestBindVarName(t *testing.T) {
	tests := []struct {
		name string
		expr ast.Expr
		want string
	}{
		{
			name: "address of ident",
			expr: &ast.UnaryExpr{Op: token.AND, X: &ast.Ident{Name: "req"}},
			want: "req",
		},
		{
			name: "plain ident",
			expr: &ast.Ident{Name: "req"},
			want: "req",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bindVarName(tt.expr)
			if got != tt.want {
				t.Errorf("bindVarName() = %q, want %q", got, tt.want)
			}
		})
	}
}
