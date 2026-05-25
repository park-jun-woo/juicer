//ff:func feature=scan type=extract control=sequence
//ff:what TestIsGinH_NilType 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestIsGinH_NilType(t *testing.T) {
	comp := &ast.CompositeLit{}
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	if isGinH(comp, info) {
		t.Fatal("expected false for nil type")
	}
}
