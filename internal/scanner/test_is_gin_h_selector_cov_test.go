//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinH_SelectorCov 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestIsGinH_SelectorCov(t *testing.T) {
	comp := &ast.CompositeLit{
		Type: &ast.SelectorExpr{X: &ast.Ident{Name: "gin"}, Sel: &ast.Ident{Name: "H"}},
	}
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	if !isGinH(comp, info) {
		t.Fatal("expected true for gin.H")
	}
}
