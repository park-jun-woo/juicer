//ff:func feature=scan type=test control=sequence
//ff:what TestProcessAssign_MiddlewareInherit 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestProcessAssign_MiddlewareInherit(t *testing.T) {
	routers := map[string]*routerInfo{
		"app": {prefix: "", middleware: []string{"logger"}},
	}
	stmt := &ast.AssignStmt{
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

	processAssign(stmt, "fiber", routers)
	ri := routers["api"]
	if len(ri.middleware) != 1 || ri.middleware[0] != "logger" {
		t.Fatalf("expected [logger], got %v", ri.middleware)
	}
}
