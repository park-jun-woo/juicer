//ff:func feature=scan type=extract control=sequence
//ff:what TestAnalyzeExpr_SelectorExprNoSelection 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestAnalyzeExpr_SelectorExprNoSelection(t *testing.T) {
	ep := &scanner.Endpoint{}
	sel := &ast.SelectorExpr{
		X:   &ast.Ident{Name: "h"},
		Sel: &ast.Ident{Name: "Method"},
	}
	info := &types.Info{
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Uses:       make(map[*ast.Ident]types.Object),
	}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	analyzeExpr(ep, sel, info, idx)
}
