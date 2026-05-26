//ff:func feature=scan type=extract control=sequence
//ff:what TestScanBody_NilBody 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestScanBody_NilBody(t *testing.T) {
	ep := &scanner.Endpoint{}
	info := &types.Info{}
	idx := &funcIndex{
		byPos: make(map[token.Pos]*ast.FuncDecl),
		info:  make(map[token.Pos]*types.Info),
	}
	scanBody(ep, nil, "c", info, idx, "handler")
	// Should return immediately
}
