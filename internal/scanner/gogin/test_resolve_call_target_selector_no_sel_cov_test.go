//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallTarget_SelectorNoSelCov 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveCallTarget_SelectorNoSelCov(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "h"}, Sel: &ast.Ident{Name: "Method"}}
	call := &ast.CallExpr{Fun: sel}
	info := &types.Info{
		Uses:       make(map[*ast.Ident]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	pos := resolveCallTarget(call, info)
	if pos != token.NoPos {
		t.Fatal("expected NoPos")
	}
}
