//ff:func feature=scan type=extract control=sequence
//ff:what TestCheckOneDepthCall_NoCtxArg 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestCheckOneDepthCall_NoCtxArg(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := &ast.CallExpr{
		Fun:  &ast.Ident{Name: "someFunc"},
		Args: []ast.Expr{&ast.BasicLit{Kind: token.INT, Value: "1"}},
	}
	info := &types.Info{
		Uses: make(map[*ast.Ident]types.Object),
	}
	idx := &funcIndex{
		byPos: make(map[token.Pos]*ast.FuncDecl),
		info:  make(map[token.Pos]*types.Info),
	}
	checkOneDepthCall(ep, call, "c", info, idx)
	// Should return since no arg matches ctxName
}
