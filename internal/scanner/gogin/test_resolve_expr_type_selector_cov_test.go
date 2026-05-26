//ff:func feature=scan type=test control=sequence
//ff:what TestResolveExprType_SelectorCov 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveExprType_SelectorCov(t *testing.T) {
	selIdent := &ast.Ident{Name: "Field"}
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "pkg"}, Sel: selIdent}
	obj := types.NewVar(0, nil, "Field", types.Typ[types.Int])
	info := &types.Info{
		Uses:  map[*ast.Ident]types.Object{selIdent: obj},
		Defs:  make(map[*ast.Ident]types.Object),
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	resolveExprType(sel, info)
}
