//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what TestIsGinInit 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestIsGinInit(t *testing.T) {
	tests := []struct {
		name  string
		sel   *ast.SelectorExpr
		alias string
		want  bool
	}{
		{
			name:  "gin.Default",
			sel:   &ast.SelectorExpr{X: &ast.Ident{Name: "gin"}, Sel: &ast.Ident{Name: "Default"}},
			alias: "gin",
			want:  true,
		},
		{
			name:  "gin.New",
			sel:   &ast.SelectorExpr{X: &ast.Ident{Name: "gin"}, Sel: &ast.Ident{Name: "New"}},
			alias: "gin",
			want:  true,
		},
		{
			name:  "wrong method",
			sel:   &ast.SelectorExpr{X: &ast.Ident{Name: "gin"}, Sel: &ast.Ident{Name: "Other"}},
			alias: "gin",
			want:  false,
		},
		{
			name:  "non-ident X",
			sel:   &ast.SelectorExpr{X: &ast.CompositeLit{}, Sel: &ast.Ident{Name: "Default"}},
			alias: "gin",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isGinInit(tt.sel, tt.alias)
			if got != tt.want {
				t.Errorf("isGinInit() = %v, want %v", got, tt.want)
			}
		})
	}
}
