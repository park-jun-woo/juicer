//ff:func feature=scan type=test control=sequence
//ff:what TestAnalyzeExpr_FuncLitWithCtx 테스트
package gogin

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestAnalyzeExpr_FuncLitWithCtx(t *testing.T) {
	ep := &scanner.Endpoint{}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}

	analyzeExpr(ep, ginCtxFuncLit("c"), nil, idx)
}
