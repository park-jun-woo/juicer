package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestLookupFunc_Found(t *testing.T) {
	pos := token.Pos(42)
	decl := &ast.FuncDecl{Name: &ast.Ident{Name: "test"}}
	info := &types.Info{}
	idx := &funcIndex{
		byPos: map[token.Pos]*ast.FuncDecl{pos: decl},
		info:  map[token.Pos]*types.Info{pos: info},
	}
	fn, inf := lookupFunc(pos, idx)
	if fn != decl {
		t.Fatal("expected decl")
	}
	if inf != info {
		t.Fatal("expected info")
	}
}

func TestLookupFunc_Missing(t *testing.T) {
	idx := &funcIndex{
		byPos: make(map[token.Pos]*ast.FuncDecl),
		info:  make(map[token.Pos]*types.Info),
	}
	fn, inf := lookupFunc(token.Pos(99), idx)
	if fn != nil || inf != nil {
		t.Fatal("expected nil")
	}
}
