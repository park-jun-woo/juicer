//ff:func feature=scan type=test control=sequence
//ff:what TestAnalyzeExpr_SelectorExprNoSelectionCov 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestAnalyzeExpr_SelectorExprNoSelectionCov(t *testing.T) {
	ep := &Endpoint{}
	sel := &ast.SelectorExpr{
		X:   &ast.Ident{Name: "h"},
		Sel: &ast.Ident{Name: "Method"},
	}
	info := &types.Info{
		Uses:       make(map[*ast.Ident]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	analyzeExpr(ep, sel, info, idx)
}
