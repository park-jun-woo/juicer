//ff:func feature=scan type=extract control=sequence
//ff:what TestCheckOneDepthCall_WithCtxArg_NoFunc 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestCheckOneDepthCall_WithCtxArg_NoFunc(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{
		Fun:  &ast.Ident{Name: "someFunc"},
		Args: []ast.Expr{&ast.Ident{Name: "c"}},
	}
	info := &types.Info{
		Uses:       make(map[*ast.Ident]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	idx := &funcIndex{
		byPos: make(map[token.Pos]*ast.FuncDecl),
		info:  make(map[token.Pos]*types.Info),
	}
	checkOneDepthCall(ep, call, "c", info, idx)
	// Should return since func not found in Uses
}
