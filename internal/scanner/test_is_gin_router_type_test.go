//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what TestIsGinRouterType 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestIsGinRouterType(t *testing.T) {
	tests := []struct {
		name  string
		expr  ast.Expr
		alias string
		want  bool
	}{
		{
			name: "*gin.Engine",
			expr: &ast.StarExpr{
				X: &ast.SelectorExpr{X: &ast.Ident{Name: "gin"}, Sel: &ast.Ident{Name: "Engine"}},
			},
			alias: "gin",
			want:  true,
		},
		{
			name: "*gin.RouterGroup",
			expr: &ast.StarExpr{
				X: &ast.SelectorExpr{X: &ast.Ident{Name: "gin"}, Sel: &ast.Ident{Name: "RouterGroup"}},
			},
			alias: "gin",
			want:  true,
		},
		{
			name:  "not a selector",
			expr:  &ast.Ident{Name: "Engine"},
			alias: "gin",
			want:  false,
		},
		{
			name: "wrong package",
			expr: &ast.StarExpr{
				X: &ast.SelectorExpr{X: &ast.Ident{Name: "other"}, Sel: &ast.Ident{Name: "Engine"}},
			},
			alias: "gin",
			want:  false,
		},
		{
			name: "wrong type",
			expr: &ast.StarExpr{
				X: &ast.SelectorExpr{X: &ast.Ident{Name: "gin"}, Sel: &ast.Ident{Name: "Context"}},
			},
			alias: "gin",
			want:  false,
		},
		{
			name:  "selector X not ident",
			expr:  &ast.SelectorExpr{X: &ast.CompositeLit{}, Sel: &ast.Ident{Name: "Engine"}},
			alias: "gin",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isGinRouterType(tt.expr, tt.alias)
			if got != tt.want {
				t.Errorf("isGinRouterType() = %v, want %v", got, tt.want)
			}
		})
	}
}
