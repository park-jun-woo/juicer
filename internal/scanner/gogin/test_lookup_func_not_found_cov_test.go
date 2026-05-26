//ff:func feature=scan type=test control=sequence
//ff:what TestLookupFunc_NotFoundCov 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestLookupFunc_NotFoundCov(t *testing.T) {
	idx := &funcIndex{
		byPos: make(map[token.Pos]*ast.FuncDecl),
		info:  make(map[token.Pos]*types.Info),
	}
	fn, _ := lookupFunc(token.Pos(999), idx)
	if fn != nil {
		t.Fatal("expected nil")
	}
}
