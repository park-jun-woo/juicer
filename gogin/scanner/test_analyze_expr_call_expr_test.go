//ff:func feature=scan type=extract control=sequence
//ff:what TestAnalyzeExpr_CallExpr 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestAnalyzeExpr_CallExpr(t *testing.T) {
	// CallExpr delegates to analyzeExpr on the Fun
	callExpr := &ast.CallExpr{
		Fun: &ast.Ident{Name: "handler"},
	}
	ep := &Endpoint{}
	info := &types.Info{
		Uses: make(map[*ast.Ident]types.Object),
	}
	idx := &funcIndex{
		byPos: make(map[token.Pos]*ast.FuncDecl),
		info:  make(map[token.Pos]*types.Info),
	}
	analyzeExpr(ep, callExpr, info, idx)
	// No crash expected
}
