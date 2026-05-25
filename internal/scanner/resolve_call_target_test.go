package scanner

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

func TestResolveCallTarget_Selector(t *testing.T) {
	call := &ast.CallExpr{Fun: &ast.SelectorExpr{
		X: &ast.Ident{Name: "h"}, Sel: &ast.Ident{Name: "Method"},
	}}
	info := &types.Info{
		Uses:       make(map[*ast.Ident]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	pos := resolveCallTarget(call, info)
	if pos != token.NoPos {
		t.Fatal("expected NoPos")
	}
}
