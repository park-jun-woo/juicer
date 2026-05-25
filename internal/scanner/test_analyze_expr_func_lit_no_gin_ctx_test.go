//ff:func feature=scan type=extract control=sequence
//ff:what TestAnalyzeExpr_FuncLit_NoGinCtx 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestAnalyzeExpr_FuncLit_NoGinCtx(t *testing.T) {
	// FuncLit without *gin.Context param — should return without panic
	funcLit := &ast.FuncLit{
		Type: &ast.FuncType{Params: &ast.FieldList{}},
		Body: &ast.BlockStmt{},
	}
	ep := &Endpoint{}
	idx := &funcIndex{
		byPos: make(map[token.Pos]*ast.FuncDecl),
		info:  make(map[token.Pos]*types.Info),
	}
	analyzeExpr(ep, funcLit, &types.Info{}, idx)
	// No crash expected
}
