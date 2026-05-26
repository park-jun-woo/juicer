//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinHMapType_IntMapCov 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestIsGinHMapType_IntMapCov(t *testing.T) {
	comp := &ast.CompositeLit{Type: &ast.Ident{Name: "M"}}
	mapType := types.NewMap(types.Typ[types.Int], types.Typ[types.String])
	info := &types.Info{Types: map[ast.Expr]types.TypeAndValue{
		comp: {Type: mapType},
	}}
	if isGinHMapType(comp, info) {
		t.Fatal("expected false for map[int]string")
	}
}
