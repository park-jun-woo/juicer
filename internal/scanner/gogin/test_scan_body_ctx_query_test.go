//ff:func feature=scan type=extract control=sequence
//ff:what TestScanBody_CtxQuery 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestScanBody_CtxQuery(t *testing.T) {
	body := &ast.BlockStmt{
		List: []ast.Stmt{
			&ast.ExprStmt{
				X: &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   &ast.Ident{Name: "c"},
						Sel: &ast.Ident{Name: "Query"},
					},
					Args: []ast.Expr{
						&ast.BasicLit{Kind: token.STRING, Value: `"name"`},
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

	if ep.Request == nil {
		t.Fatal("expected request to be set")
	}
	if len(ep.Request.Query) != 1 {
		t.Errorf("expected 1 query param, got %d", len(ep.Request.Query))
	}
	if ep.Request.Query[0].Name != "name" {
		t.Errorf("expected query param name 'name', got %q", ep.Request.Query[0].Name)
	}
}
