//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinHMapType_StringMapCov 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestIsGinHMapType_StringMapCov(t *testing.T) {
	comp := &ast.CompositeLit{Type: &ast.Ident{Name: "H"}}
	mapType := types.NewMap(types.Typ[types.String], types.NewInterfaceType(nil, nil))
	info := &types.Info{Types: map[ast.Expr]types.TypeAndValue{
		comp: {Type: mapType},
	}}
	if !isGinHMapType(comp, info) {
		t.Fatal("expected true for map[string]any")
	}
}
