//ff:func feature=scan type=extract control=sequence
//ff:what TestCheckOneDepthCall_SelectorExpr_WithUses 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestCheckOneDepthCall_SelectorExpr_WithUses(t *testing.T) {
	// Test with a SelectorExpr where Uses has the identifier
	ep := &Endpoint{}
	sel := &ast.SelectorExpr{
		X:   &ast.Ident{Name: "pkg"},
		Sel: &ast.Ident{Name: "Func"},
	}
	call := &ast.CallExpr{
		Fun:  sel,
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
}
