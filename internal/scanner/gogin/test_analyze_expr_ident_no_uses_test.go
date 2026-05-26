//ff:func feature=scan type=extract control=sequence
//ff:what TestAnalyzeExpr_Ident_NoUses 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestAnalyzeExpr_Ident_NoUses(t *testing.T) {
	ident := &ast.Ident{Name: "handler"}
	ep := &scanner.Endpoint{}
	info := &types.Info{
		Uses: make(map[*ast.Ident]types.Object),
	}
	idx := &funcIndex{
		byPos: make(map[token.Pos]*ast.FuncDecl),
		info:  make(map[token.Pos]*types.Info),
	}
	analyzeExpr(ep, ident, info, idx)
	// No crash expected
}
