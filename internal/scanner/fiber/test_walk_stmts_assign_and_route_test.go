//ff:func feature=scan type=test control=sequence
//ff:what TestWalkStmts_AssignAndRoute 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestWalkStmts_AssignAndRoute(t *testing.T) {
	fset := token.NewFileSet()
	fset.AddFile("main.go", 1, 100)

	routers := map[string]*routerInfo{
		"app": {prefix: ""},
	}

	// Simulates: api := app.Group("/api")
	assignStmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "api"}},
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "app"},
					Sel: &ast.Ident{Name: "Group"},
				},
				Args: []ast.Expr{
					&ast.BasicLit{Kind: token.STRING, Value: `"/api"`},
				},
			},
		},
		Tok: token.DEFINE,
	}

	// Simulates: api.Get("/users", listUsers)
	routeStmt := &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "api"},
				Sel: &ast.Ident{Name: "Get"},
			},
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: `"/users"`},
				&ast.Ident{Name: "listUsers"},
			},
		},
	}

	var eps []scanner.Endpoint
	hmap := map[int][]ast.Expr{}
	walkStmts([]ast.Stmt{assignStmt, routeStmt}, "fiber", "main.go", fset, routers, &eps, hmap)

	if len(eps) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(eps))
	}
	if eps[0].Path != "/api/users" {
		t.Fatalf("expected /api/users, got %s", eps[0].Path)
	}
	if eps[0].Method != "GET" {
		t.Fatalf("expected GET, got %s", eps[0].Method)
	}
}
