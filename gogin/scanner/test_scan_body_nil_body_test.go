//ff:func feature=scan type=extract control=sequence
//ff:what TestScanBody_NilBody 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestScanBody_NilBody(t *testing.T) {
	ep := &Endpoint{}
	info := &types.Info{}
	idx := &funcIndex{
		byPos: make(map[token.Pos]*ast.FuncDecl),
		info:  make(map[token.Pos]*types.Info),
	}
	scanBody(ep, nil, "c", info, idx, "handler")
	// Should return immediately
}
