//ff:func feature=scan type=test control=sequence
//ff:what TestWalkStmts_Middleware 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestWalkStmts_Middleware(t *testing.T) {
	fset := token.NewFileSet()
	fset.AddFile("main.go", 1, 100)

	routers := map[string]*routerInfo{
		"app": {prefix: ""},
	}

	// Simulates: app.Use(logger)
	useStmt := &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "app"},
				Sel: &ast.Ident{Name: "Use"},
			},
			Args: []ast.Expr{
				&ast.Ident{Name: "logger"},
			},
		},
	}

	// Simulates: app.Get("/health", healthCheck)
	routeStmt := &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "app"},
				Sel: &ast.Ident{Name: "Get"},
			},
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: `"/health"`},
				&ast.Ident{Name: "healthCheck"},
			},
		},
	}

	var eps []scanner.Endpoint
	hmap := map[int][]ast.Expr{}
	walkStmts([]ast.Stmt{useStmt, routeStmt}, "fiber", "main.go", fset, routers, &eps, hmap)

	if len(eps) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(eps))
	}
	if len(eps[0].Middleware) != 1 || eps[0].Middleware[0] != "logger" {
		t.Fatalf("expected [logger] middleware, got %v", eps[0].Middleware)
	}
}
