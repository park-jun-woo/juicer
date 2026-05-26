//ff:func feature=scan type=test control=sequence
//ff:what TestCheckOneDepthCall_NoCtxInArgs 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestCheckOneDepthCall_NoCtxInArgs(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{
		Fun:  &ast.Ident{Name: "f"},
		Args: []ast.Expr{&ast.Ident{Name: "other"}},
	}
	info := &types.Info{Uses: make(map[*ast.Ident]types.Object)}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	checkOneDepthCall(ep, call, "c", info, idx)
}
