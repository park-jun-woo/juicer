//ff:func feature=scan type=extract control=sequence
//ff:what TestScanBody_WrongReceiver 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestScanBody_WrongReceiver(t *testing.T) {
	// A selector call but not on ctxName
	body := &ast.BlockStmt{
		List: []ast.Stmt{
			&ast.ExprStmt{
				X: &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   &ast.Ident{Name: "other"},
						Sel: &ast.Ident{Name: "Method"},
					},
				},
			},
		},
	}
	ep := &scanner.Endpoint{}
	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Uses:       make(map[*ast.Ident]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	idx := &funcIndex{
		byPos: make(map[token.Pos]*ast.FuncDecl),
		info:  make(map[token.Pos]*types.Info),
	}
	scanBody(ep, body, "c", info, idx, "handler")
}
