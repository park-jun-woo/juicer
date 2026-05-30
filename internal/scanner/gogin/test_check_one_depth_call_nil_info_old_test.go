//ff:func feature=scan type=test control=sequence
//ff:what TestCheckOneDepthCall_NilInfoOld 테스트
package gogin

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestCheckOneDepthCall_NilInfoOld(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := &ast.CallExpr{Fun: &ast.Ident{Name: "f"}}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	checkOneDepthCall(ep, call, "c", nil, idx)

	info := &types.Info{Uses: make(map[*ast.Ident]types.Object)}
	call2 := &ast.CallExpr{
		Fun:  &ast.Ident{Name: "f"},
		Args: []ast.Expr{&ast.Ident{Name: "other"}},
	}
	checkOneDepthCall(ep, call2, "c", info, idx)

	call3 := &ast.CallExpr{
		Fun:  &ast.Ident{Name: "f"},
		Args: []ast.Expr{&ast.Ident{Name: "c"}},
	}
	checkOneDepthCall(ep, call3, "c", info, idx)
}
