//ff:func feature=scan type=extract control=sequence
//ff:what TestLookupFunc 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestLookupFunc(t *testing.T) {
	fnDecl := &ast.FuncDecl{
		Name: &ast.Ident{Name: "test", NamePos: 42},
		Body: &ast.BlockStmt{},
	}
	info := &types.Info{}
	idx := &funcIndex{
		byPos: map[token.Pos]*ast.FuncDecl{42: fnDecl},
		info:  map[token.Pos]*types.Info{42: info},
	}

	fn, fnInfo := lookupFunc(42, idx)
	if fn != fnDecl {
		t.Error("expected to find the function")
	}
	if fnInfo != info {
		t.Error("expected to find the info")
	}
}
