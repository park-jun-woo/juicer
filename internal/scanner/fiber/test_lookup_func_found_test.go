//ff:func feature=scan type=test control=sequence
//ff:what TestLookupFunc_Found 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestLookupFunc_Found(t *testing.T) {
	fn := &ast.FuncDecl{Name: ast.NewIdent("H")}
	info := &types.Info{}
	idx := &funcIndex{
		byPos: map[token.Pos]*ast.FuncDecl{token.Pos(5): fn},
		info:  map[token.Pos]*types.Info{token.Pos(5): info},
	}
	gotFn, gotInfo := lookupFunc(token.Pos(5), idx)
	if gotFn != fn || gotInfo != info {
		t.Fatalf("lookupFunc mismatch: %v %v", gotFn, gotInfo)
	}
}
