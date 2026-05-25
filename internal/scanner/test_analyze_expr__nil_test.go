//ff:func feature=scan type=extract control=sequence
//ff:what TestAnalyzeExpr_Nil 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestAnalyzeExpr_Nil(t *testing.T) {
	ep := &Endpoint{}
	info := &types.Info{Uses: make(map[*ast.Ident]types.Object)}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	analyzeExpr(ep, nil, info, idx)
}
