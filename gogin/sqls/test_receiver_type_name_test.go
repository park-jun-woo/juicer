//ff:func feature=sql type=parse control=iteration dimension=1
//ff:what TestReceiverTypeName 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestReceiverTypeName(t *testing.T) {
	tests := []struct {
		name string
		expr ast.Expr
		want string
	}{
		{"star", &ast.StarExpr{X: &ast.Ident{Name: "UserRepo"}}, "UserRepo"},
		{"ident", &ast.Ident{Name: "UserRepo"}, "UserRepo"},
		{"star non-ident", &ast.StarExpr{X: &ast.SelectorExpr{X: &ast.Ident{Name: "pkg"}, Sel: &ast.Ident{Name: "Type"}}}, ""},
		{"unknown", &ast.CompositeLit{}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := receiverTypeName(tt.expr)
			if got != tt.want {
				t.Errorf("receiverTypeName() = %q, want %q", got, tt.want)
			}
		})
	}
}
