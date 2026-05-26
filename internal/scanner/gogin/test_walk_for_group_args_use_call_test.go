//ff:func feature=scan type=test control=sequence
//ff:what walkForGroupArgs에서 Use() 호출이 routers 맵의 middleware에 반영되는지 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestWalkForGroupArgs_UseCall(t *testing.T) {
	fset := token.NewFileSet()
	routers := map[string]*routerInfo{
		"authGroup": {prefix: "/api", middleware: []string{}},
	}
	ctx := &groupArgCtx{
		ginAlias:  "gin",
		routers:   routers,
		fset:      fset,
		idx:       &funcIndex{},
		root:      "/tmp",
		endpoints: nil,
		epIndex:   map[struct{ file string; line int }]int{},
	}

	// authGroup.Use(middleware.Auth())
	stmts := []ast.Stmt{
		&ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "authGroup"},
					Sel: &ast.Ident{Name: "Use"},
				},
				Args: []ast.Expr{
					&ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   &ast.Ident{Name: "middleware"},
							Sel: &ast.Ident{Name: "Auth"},
						},
					},
				},
			},
		},
	}
	walkForGroupArgs(stmts, ctx)

	if len(routers["authGroup"].middleware) != 1 {
		t.Fatalf("expected 1 middleware, got %d", len(routers["authGroup"].middleware))
	}
	if routers["authGroup"].middleware[0] != "middleware.Auth()" {
		t.Errorf("expected middleware.Auth(), got %q", routers["authGroup"].middleware[0])
	}
}
