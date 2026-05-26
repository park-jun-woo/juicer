//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallTarget_Ident 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveCallTarget_Ident(t *testing.T) {
	call := &ast.CallExpr{Fun: &ast.Ident{Name: "f"}}
	info := &types.Info{
		Uses:       make(map[*ast.Ident]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	pos := resolveCallTarget(call, info)
	if pos != token.NoPos {
		t.Fatal("expected NoPos")
	}
}

