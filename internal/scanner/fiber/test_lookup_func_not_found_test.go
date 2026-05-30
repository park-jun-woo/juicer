//ff:func feature=scan type=test control=sequence
//ff:what TestLookupFunc_NotFound 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestLookupFunc_NotFound(t *testing.T) {
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}, info: map[token.Pos]*types.Info{}}
	fn, info := lookupFunc(token.Pos(99), idx)
	if fn != nil || info != nil {
		t.Fatalf("expected nil,nil for missing pos, got %v %v", fn, info)
	}
}
