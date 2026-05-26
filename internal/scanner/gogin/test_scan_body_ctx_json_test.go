//ff:func feature=scan type=extract control=sequence
//ff:what TestScanBody_CtxJSON 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestScanBody_CtxJSON(t *testing.T) {
	body := &ast.BlockStmt{
		List: []ast.Stmt{
			&ast.ExprStmt{
				X: &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   &ast.Ident{Name: "c"},
						Sel: &ast.Ident{Name: "JSON"},
					},
					Args: []ast.Expr{
						&ast.BasicLit{Kind: token.INT, Value: "200"},
						&ast.Ident{Name: "result"},
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

	if len(ep.Responses) != 1 {
		t.Fatalf("expected 1 response, got %d", len(ep.Responses))
	}
	if ep.Responses[0].Status != "200" {
		t.Errorf("expected status 200, got %q", ep.Responses[0].Status)
	}
	if ep.Responses[0].Kind != "json" {
		t.Errorf("expected kind 'json', got %q", ep.Responses[0].Kind)
	}
}
