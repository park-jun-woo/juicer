//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinH_NonGinCov 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestIsGinH_NonGinCov(t *testing.T) {
	comp := &ast.CompositeLit{Type: &ast.Ident{Name: "Foo"}}
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	if isGinH(comp, info) {
		t.Fatal("expected false for non-gin type")
	}
}
