//ff:func feature=scan type=extract control=sequence
//ff:what TestCheckOneDepthCall_NilInfo 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestCheckOneDepthCall_NilInfo(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := &ast.CallExpr{
		Fun: &ast.Ident{Name: "someFunc"},
	}
	idx := &funcIndex{
		byPos: make(map[token.Pos]*ast.FuncDecl),
		info:  make(map[token.Pos]*types.Info),
	}
	checkOneDepthCall(ep, call, "c", nil, idx)
	// Should return immediately since info is nil
	// Should return immediately
}
