//ff:func feature=scan type=test control=sequence
//ff:what TestWalkStmts_RangeStmt — range 내부 라우트 감지 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestWalkStmts_RangeStmt(t *testing.T) {
	routeCall := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "r"},
			Sel: &ast.Ident{Name: "GET"},
		},
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.STRING, Value: `"/in-range"`},
			&ast.Ident{Name: "handler"},
		},
	}
	rangeStmt := &ast.RangeStmt{
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ExprStmt{X: routeCall},
			},
		},
	}

	routers := map[string]*routerInfo{"r": {}}
	var out []scanner.Endpoint
	fset := token.NewFileSet()
	walkStmts([]ast.Stmt{rangeStmt}, "gin", "test.go", fset, routers, &out, map[int][]ast.Expr{})
	if len(out) != 1 {
		t.Fatalf("expected 1 endpoint from range, got %d", len(out))
	}
	if out[0].Path != "/in-range" {
		t.Errorf("expected /in-range, got %s", out[0].Path)
	}
}
