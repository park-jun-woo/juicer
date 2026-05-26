//ff:func feature=scan type=test control=sequence
//ff:what TestLookupFunc_Found 테스트
package gogin

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

