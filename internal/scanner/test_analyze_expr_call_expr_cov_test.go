//ff:func feature=scan type=test control=sequence
//ff:what TestAnalyzeExpr_CallExprCov 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestAnalyzeExpr_CallExprCov(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{
		Fun: &ast.Ident{Name: "makeHandler"},
	}
	info := &types.Info{
		Uses:       make(map[*ast.Ident]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	analyzeExpr(ep, call, info, idx)
}
