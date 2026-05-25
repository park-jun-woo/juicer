//ff:func feature=scan type=extract control=sequence
//ff:what TestAnalyzeExpr_CallExprCase 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestAnalyzeExpr_CallExprCase(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{
		Fun: &ast.Ident{Name: "handler"},
	}
	info := &types.Info{Uses: make(map[*ast.Ident]types.Object)}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	analyzeExpr(ep, call, info, idx)
}
