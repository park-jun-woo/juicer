//ff:func feature=scan type=extract control=sequence
//ff:what TestScanBody_EmptyBody 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestScanBody_EmptyBody(t *testing.T) {
	ep := &Endpoint{}
	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Uses:       make(map[*ast.Ident]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	idx := &funcIndex{
		byPos: make(map[token.Pos]*ast.FuncDecl),
		info:  make(map[token.Pos]*types.Info),
	}
	scanBody(ep, &ast.BlockStmt{}, "c", info, idx, "handler")
	// No crash expected
}
