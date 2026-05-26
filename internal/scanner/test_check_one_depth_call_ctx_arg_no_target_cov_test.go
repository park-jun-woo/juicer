//ff:func feature=scan type=test control=sequence
//ff:what TestCheckOneDepthCall_CtxArgNoTargetCov 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestCheckOneDepthCall_CtxArgNoTargetCov(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{
		Fun:  &ast.Ident{Name: "doSomething"},
		Args: []ast.Expr{&ast.Ident{Name: "c"}},
	}
	info := &types.Info{
		Uses:       make(map[*ast.Ident]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	checkOneDepthCall(ep, call, "c", info, idx)
}
