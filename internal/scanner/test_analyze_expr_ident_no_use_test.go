//ff:func feature=scan type=test control=sequence
//ff:what TestAnalyzeExpr_IdentNoUse 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestAnalyzeExpr_IdentNoUse(t *testing.T) {
	ep := &Endpoint{}
	ident := &ast.Ident{Name: "handler"}
	info := &types.Info{Uses: make(map[*ast.Ident]types.Object)}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	analyzeExpr(ep, ident, info, idx)
}
