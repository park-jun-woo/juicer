//ff:func feature=scan type=test control=sequence
//ff:what TestProcessAssign_NestedGroup 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestProcessAssign_NestedGroup(t *testing.T) {
	routers := map[string]*routerInfo{
		"api": {prefix: "/api"},
	}
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "users"}},
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "api"},
					Sel: &ast.Ident{Name: "Group"},
				},
				Args: []ast.Expr{
					&ast.BasicLit{Kind: token.STRING, Value: `"/users"`},
				},
			},
		},
		Tok: token.DEFINE,
	}

	processAssign(stmt, "fiber", routers)
	ri, ok := routers["users"]
	if !ok {
		t.Fatal("expected users in routers")
	}
	if ri.prefix != "/api/users" {
		t.Fatalf("expected /api/users, got %s", ri.prefix)
	}
}
