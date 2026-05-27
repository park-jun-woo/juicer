//ff:func feature=scan type=test control=sequence
//ff:what TestCheckOneDepthCall_WithCtxNoTarget 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestCheckOneDepthCall_WithCtxNoTarget(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := &ast.CallExpr{
		Fun:  &ast.Ident{Name: "f"},
		Args: []ast.Expr{&ast.Ident{Name: "c"}},
	}
	info := &types.Info{Uses: make(map[*ast.Ident]types.Object)}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	checkOneDepthCall(ep, call, "c", info, idx)
}
