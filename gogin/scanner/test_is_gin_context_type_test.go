//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what TestIsGinContextType 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestIsGinContextType(t *testing.T) {
	tests := []struct {
		name string
		expr ast.Expr
		want bool
	}{
		{
			name: "valid *gin.Context",
			expr: &ast.StarExpr{
				X: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "gin"},
					Sel: &ast.Ident{Name: "Context"},
				},
			},
			want: true,
		},
		{
			name: "wrong selector name",
			expr: &ast.StarExpr{
				X: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "gin"},
					Sel: &ast.Ident{Name: "Engine"},
				},
			},
			want: false,
		},
		{
			name: "not a star expr",
			expr: &ast.Ident{Name: "int"},
			want: false,
		},
		{
			name: "star but not selector",
			expr: &ast.StarExpr{X: &ast.Ident{Name: "int"}},
			want: false,
		},
		{
			name: "selector but X not ident",
			expr: &ast.StarExpr{
				X: &ast.SelectorExpr{
					X:   &ast.CompositeLit{},
					Sel: &ast.Ident{Name: "Context"},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isGinContextType(tt.expr)
			if got != tt.want {
				t.Errorf("isGinContextType() = %v, want %v", got, tt.want)
			}
		})
	}
}
