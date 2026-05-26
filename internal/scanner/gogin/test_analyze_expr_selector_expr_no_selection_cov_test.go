//ff:func feature=scan type=test control=sequence
//ff:what TestAnalyzeExpr_SelectorExprNoSelectionCov 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestAnalyzeExpr_SelectorExprNoSelectionCov(t *testing.T) {
	ep := &scanner.Endpoint{}
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
