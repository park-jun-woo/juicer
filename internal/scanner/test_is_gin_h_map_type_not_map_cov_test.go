//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinHMapType_NotMapCov 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestIsGinHMapType_NotMapCov(t *testing.T) {
	comp := &ast.CompositeLit{Type: &ast.Ident{Name: "X"}}
	info := &types.Info{Types: map[ast.Expr]types.TypeAndValue{
		comp: {Type: types.Typ[types.Int]},
	}}
	if isGinHMapType(comp, info) {
		t.Fatal("expected false for non-map")
	}
}
