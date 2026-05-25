//ff:func feature=scan type=extract control=sequence
//ff:what TestLookupFunc_NotFound 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestLookupFunc_NotFound(t *testing.T) {
	idx := &funcIndex{
		byPos: make(map[token.Pos]*ast.FuncDecl),
		info:  make(map[token.Pos]*types.Info),
	}

	fn, fnInfo := lookupFunc(99, idx)
	if fn != nil {
		t.Error("expected nil for non-existent pos")
	}
	if fnInfo != nil {
		t.Error("expected nil info for non-existent pos")
	}
}
