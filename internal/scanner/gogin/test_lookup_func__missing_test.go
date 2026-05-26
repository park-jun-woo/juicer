//ff:func feature=scan type=extract control=sequence
//ff:what TestLookupFunc_Missing 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

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
