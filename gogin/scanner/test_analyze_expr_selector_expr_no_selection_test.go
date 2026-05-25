//ff:func feature=scan type=extract control=sequence
//ff:what TestAnalyzeExpr_SelectorExpr_NoSelection 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestAnalyzeExpr_SelectorExpr_NoSelection(t *testing.T) {
	sel := &ast.SelectorExpr{
		X:   &ast.Ident{Name: "h"},
		Sel: &ast.Ident{Name: "Method"},
	}
	ep := &Endpoint{}
	info := &types.Info{
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	idx := &funcIndex{
		byPos: make(map[token.Pos]*ast.FuncDecl),
		info:  make(map[token.Pos]*types.Info),
	}
	analyzeExpr(ep, sel, info, idx)
	// No crash expected
}
