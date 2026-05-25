//ff:func feature=scan type=extract control=sequence
//ff:what TestScanBody_CtxParam 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestScanBody_CtxParam(t *testing.T) {
	body := &ast.BlockStmt{
		List: []ast.Stmt{
			&ast.ExprStmt{
				X: &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   &ast.Ident{Name: "c"},
						Sel: &ast.Ident{Name: "Param"},
					},
					Args: []ast.Expr{
						&ast.BasicLit{Kind: token.STRING, Value: `"id"`},
					},
				},
			},
		},
	}
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
	scanBody(ep, body, "c", info, idx, "handler")

	if ep.Request == nil {
		t.Fatal("expected request to be set")
	}
	if len(ep.Request.PathParams) != 1 {
		t.Errorf("expected 1 path param, got %d", len(ep.Request.PathParams))
	}
}
