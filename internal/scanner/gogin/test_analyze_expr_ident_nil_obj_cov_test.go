//ff:func feature=scan type=test control=sequence
//ff:what TestAnalyzeExpr_IdentNilObjCov 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestAnalyzeExpr_IdentNilObjCov(t *testing.T) {
	ep := &scanner.Endpoint{}
	ident := &ast.Ident{Name: "handler"}
	info := &types.Info{
		Uses:       make(map[*ast.Ident]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	analyzeExpr(ep, ident, info, idx)
}
